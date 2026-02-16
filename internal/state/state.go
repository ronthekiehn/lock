package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronthekiehn/lock/internal/hosts"
)

const Version = 1

type DomainRecord struct {
	LockedAt string `json:"locked_at,omitempty"`
	Note     string `json:"note,omitempty"`
}

type File struct {
	StateVersion int                     `json:"state_version"`
	Domains      map[string]DomainRecord `json:"domains"`
}

func Path() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "Library", "Application Support", "lock", "state.json")
}

func SyncAfterLock(touched []string, note string, now time.Time, hostRecords []hosts.LockRecord) error {
	statePath := Path()
	if statePath == "" {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(statePath), 0o755); err != nil {
		return err
	}

	raw, current, validCurrent := loadState(statePath)
	nowValue := now.UTC().Format(time.RFC3339)

	if validCurrent {
		if current.Domains == nil {
			current.Domains = map[string]DomainRecord{}
		}
		for _, domain := range touched {
			record := current.Domains[domain]
			record.LockedAt = nowValue
			if note != "" {
				record.Note = note
			} else {
				record.Note = ""
			}
			current.Domains[domain] = record
		}
		current.StateVersion = Version
		return writeAtomic(statePath, current)
	}

	hostNotes := make(map[string]string)
	for _, rec := range hostRecords {
		hostNotes[rec.Domain] = rec.Note
	}

	touchedSet := make(map[string]struct{}, len(touched))
	for _, domain := range touched {
		touchedSet[domain] = struct{}{}
	}

	next := File{
		StateVersion: Version,
		Domains:      make(map[string]DomainRecord),
	}

	for _, domain := range hosts.ActiveDomains(hostRecords) {
		record := extractRecord(raw, domain)
		if hostNote := hostNotes[domain]; hostNote != "" {
			record.Note = hostNote
		}
		if _, ok := touchedSet[domain]; ok {
			record.LockedAt = nowValue
			if note != "" {
				record.Note = note
			} else {
				record.Note = ""
			}
		}
		next.Domains[domain] = record
	}

	return writeAtomic(statePath, next)
}

func StatusLines(hostRecords []hosts.LockRecord, now time.Time) []string {
	statePath := Path()
	if len(hostRecords) == 0 {
		return nil
	}

	stateDomains := map[string]DomainRecord{}
	if loaded, ok := loadValid(statePath); ok {
		stateDomains = loaded.Domains
	}

	lines := make([]string, 0, len(hostRecords))
	for _, hostRecord := range hostRecords {
		lockedFor := "unknown"
		stateRecord, exists := stateDomains[hostRecord.Domain]
		if exists {
			if since, ok := parseTimestamp(stateRecord.LockedAt); ok {
				lockedFor = formatDuration(now.UTC().Sub(since.UTC()))
			}
		}

		note := hostRecord.Note
		if exists && stateRecord.Note != "" {
			note = stateRecord.Note
		}
		safeNote := strings.ReplaceAll(note, `"`, `\"`)

		lines = append(lines, fmt.Sprintf(`%s locked_for=%s note="%s"`, hostRecord.Domain, lockedFor, safeNote))
	}

	return lines
}

func loadValid(path string) (File, bool) {
	_, loaded, valid := loadState(path)
	if !valid {
		return File{}, false
	}
	return loaded, true
}

func loadState(path string) (map[string]any, File, bool) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, File{}, false
	}

	var raw map[string]any
	if err := json.Unmarshal(content, &raw); err != nil {
		return nil, File{}, false
	}

	var parsed File
	if err := json.Unmarshal(content, &parsed); err != nil {
		return raw, File{}, false
	}
	if parsed.StateVersion != Version || parsed.Domains == nil {
		return raw, File{}, false
	}

	return raw, parsed, true
}

func extractRecord(raw map[string]any, domain string) DomainRecord {
	if raw == nil {
		return DomainRecord{}
	}

	rawDomains, ok := raw["domains"].(map[string]any)
	if !ok {
		return DomainRecord{}
	}
	rawRecord, ok := rawDomains[domain].(map[string]any)
	if !ok {
		return DomainRecord{}
	}

	record := DomainRecord{}
	if lockedAt, ok := rawRecord["locked_at"].(string); ok {
		if _, valid := parseTimestamp(lockedAt); valid {
			record.LockedAt = lockedAt
		}
	}
	if note, ok := rawRecord["note"].(string); ok && note != "" {
		record.Note = note
	}

	return record
}

func writeAtomic(path string, state File) error {
	content, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	content = append(content, '\n')

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".state.*.json")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()

	writeErr := func() error {
		defer tmp.Close()
		if _, err := tmp.Write(content); err != nil {
			return err
		}
		if err := tmp.Sync(); err != nil {
			return err
		}
		return tmp.Close()
	}()
	if writeErr != nil {
		_ = os.Remove(tmpPath)
		return writeErr
	}

	if err := os.Rename(tmpPath, path); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	return nil
}

func parseTimestamp(value string) (time.Time, bool) {
	if value == "" {
		return time.Time{}, false
	}
	ts, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Time{}, false
	}
	return ts, true
}

func formatDuration(duration time.Duration) string {
	if duration < 0 {
		duration = 0
	}
	totalSeconds := int64(duration.Seconds())
	days := totalSeconds / 86400
	totalSeconds = totalSeconds % 86400
	hours := totalSeconds / 3600
	totalSeconds = totalSeconds % 3600
	minutes := totalSeconds / 60

	var b strings.Builder
	if days > 0 {
		b.WriteString(fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		b.WriteString(fmt.Sprintf("%dh", hours))
	}
	b.WriteString(fmt.Sprintf("%dm", minutes))
	return b.String()
}

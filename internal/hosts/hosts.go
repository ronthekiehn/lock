package hosts

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ronthekiehn/lock/internal/domain"
)

const DefaultPath = "/etc/hosts"

type LockRecord struct {
	Domain string
	Note   string
}

func LockRecords(path string) ([]LockRecord, error) {
	file, err := os.Open(path)
	if err != nil {
		// Shell script silences read errors and treats them as no records.
		return nil, nil
	}
	defer file.Close()

	return parseLockRecords(file)
}

func parseLockRecords(reader io.Reader) ([]LockRecord, error) {
	scanner := bufio.NewScanner(reader)
	currentDomain := ""
	currentNote := ""
	seen := make(map[string]struct{})
	var records []LockRecord

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLeft := strings.TrimLeft(line, " \t")

		switch {
		case strings.HasPrefix(trimmedLeft, "# lock:"):
			currentDomain = strings.TrimSpace(strings.TrimPrefix(trimmedLeft, "# lock:"))
			currentNote = ""
			continue
		case strings.HasPrefix(trimmedLeft, "# note:"):
			if currentDomain != "" {
				currentNote = strings.TrimSpace(strings.TrimPrefix(trimmedLeft, "# note:"))
			}
			continue
		case strings.HasPrefix(trimmedLeft, "0.0.0.0"):
			if currentDomain == "" {
				continue
			}
			fields := strings.Fields(trimmedLeft)
			if len(fields) < 2 {
				continue
			}
			host := fields[1]
			normalizedDomain := domain.Normalize(currentDomain)
			normalizedHost := domain.Normalize(host)
			if !domain.IsValid(normalizedDomain) {
				continue
			}
			if normalizedHost != normalizedDomain && normalizedHost != "www."+normalizedDomain {
				continue
			}
			if _, exists := seen[normalizedDomain]; exists {
				continue
			}
			seen[normalizedDomain] = struct{}{}
			records = append(records, LockRecord{Domain: normalizedDomain, Note: currentNote})
			continue
		case strings.HasPrefix(trimmedLeft, "#"):
			continue
		case strings.TrimSpace(line) == "":
			// Keep current lock block context across blank lines.
			continue
		default:
			currentDomain = ""
			currentNote = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func ActiveDomains(records []LockRecord) []string {
	out := make([]string, 0, len(records))
	for _, record := range records {
		out = append(out, record.Domain)
	}
	return out
}

func IsDomainLocked(path, lockDomain string) (bool, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, nil
	}
	needle := []byte("0.0.0.0 " + lockDomain)
	return bytes.Contains(content, needle), nil
}

func AppendLockEntry(path, lockDomain string, includeNote bool, sanitizedNote string) error {
	var b strings.Builder
	b.WriteString("# lock: ")
	b.WriteString(lockDomain)
	b.WriteByte('\n')
	if includeNote {
		b.WriteString("# note: ")
		b.WriteString(sanitizedNote)
		b.WriteByte('\n')
	}
	b.WriteString("0.0.0.0 ")
	b.WriteString(lockDomain)
	b.WriteByte('\n')
	b.WriteString("0.0.0.0 www.")
	b.WriteString(lockDomain)
	b.WriteByte('\n')

	cmd := exec.Command("sudo", "tee", "-a", path)
	cmd.Stdin = strings.NewReader(b.String())
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

package domain

import (
	"regexp"
	"strings"
)

var domainPattern = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]*[a-z0-9])?(\.[a-z0-9]([a-z0-9-]*[a-z0-9])?)+$`)

// Normalize mirrors the existing shell behavior:
// strip scheme, strip leading www, strip path, then lowercase.
func Normalize(input string) string {
	value := input
	value = strings.TrimPrefix(value, "http://")
	value = strings.TrimPrefix(value, "https://")
	value = strings.TrimPrefix(value, "www.")
	if idx := strings.Index(value, "/"); idx >= 0 {
		value = value[:idx]
	}
	return strings.ToLower(value)
}

func IsValid(value string) bool {
	return domainPattern.MatchString(value)
}

func NormalizeAndClassify(inputs []string) (valid []string, invalid []string) {
	seen := make(map[string]struct{})
	for _, input := range inputs {
		normalized := Normalize(input)
		if normalized == "" || !IsValid(normalized) {
			invalid = append(invalid, input)
			continue
		}
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}
		valid = append(valid, normalized)
	}
	return valid, invalid
}

func SanitizeNote(note string) string {
	replaced := strings.NewReplacer("\r", " ", "\n", " ").Replace(note)
	collapsed := strings.Join(strings.Fields(replaced), " ")
	return strings.TrimSpace(collapsed)
}

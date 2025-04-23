package account

import "strings"

func ValidName(name string) bool {
	return strings.TrimSpace(name) != "" && len(name) > 3
}

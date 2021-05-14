package ula

import (
	"fmt"
	"strings"
	"unicode"
)

// Ula2 func CamelCase to snake_case
func Ula2(s string) string {
	var res []string

	for i, r := range s {
		if i == 0 {
			res = append(res, strings.ToLower(s[i:i+1]))
			continue
		}

		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			res = append(res, fmt.Sprintf("_%s", strings.ToLower(s[i:i+1])))
		} else {
			res = append(res, s[i:i+1])
		}
	}

	return strings.Join(res, "")
}

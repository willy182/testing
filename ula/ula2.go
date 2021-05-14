package ula

import (
	"fmt"
	"strings"
	"unicode"
)

// CamelToSnake func CamelCase to snake_case
func CamelToSnake(s string) string {
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

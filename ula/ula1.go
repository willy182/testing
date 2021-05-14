package ula

import (
	"fmt"
	"strings"
)

// GetNumberLetter func count occurrences
// input = occurrences
// output = o1c3u1r2e2n1s1
// output must be lower case
func GetNumberLetter(s string) string {
	var (
		concat   string
		res, tmp []string
	)

	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		if !strings.Contains(concat, str) {
			t := strings.Count(s, s[i:i+1])
			res = append(res, fmt.Sprintf("%s%d", str, t))
			tmp = append(tmp, str)
			concat = strings.Join(tmp, "")
		}
	}

	return strings.Join(res, "")
}

package common

import "strings"

func Palindrome(str string) bool {
	length := len(str)
	word := strings.ToLower(str)

	for i := 0; i < length; i++ {
		if word[i:i+1] != word[length-(i+1):length-i] {
			return false
		}
	}

	return true
}

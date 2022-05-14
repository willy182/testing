package alami

import (
	"strconv"
	"strings"
)

// NumericReverse function
// 234 = 432
func NumericReverse(n int) int {
	str := strconv.Itoa(n)

	var strArr []string

	for i := 0; i < len(str); i++ {
		strArr = append(strArr, str[len(str)-(i+1):len(str)-i])
	}

	num, _ := strconv.Atoi(strings.Join(strArr, ""))

	return num
}

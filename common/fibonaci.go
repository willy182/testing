package common

import (
	"strconv"
	"strings"
)

func Fibonaci1(limit int) string {
	var (
		arr []int
		str []string
	)

	for i := 0; i < limit; i++ {
		if i > 1 {
			var j int
			j = arr[i-2] + arr[i-1]
			arr = append(arr, j)
			str = append(str, strconv.Itoa(j))
		} else {
			arr = append(arr, i)
			str = append(str, strconv.Itoa(i))
		}
	}

	return strings.Join(str, ", ")
}

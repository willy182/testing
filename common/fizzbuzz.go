package common

import "fmt"

func FizzBuzz() string {
	var str string

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				str = fmt.Sprintf("%s%s", str, "Fizz")
			} else {
				str = fmt.Sprintf("%s%s", str, "Fizz\n")
			}
		}
		if i%5 == 0 {
			str = fmt.Sprintf("%s%s", str, "Buzz\n")
		}
		if (i%5 != 0) && (i%3 != 0) {
			str = fmt.Sprintf("%s%d\n", str, i)
		}
	}

	return str
}

package okadoc

import (
	"sort"
)

func FindMaxSum(numbers []int) int {
	length := len(numbers)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers[length-1] + numbers[length-2]
}

func FindMaxSumSortRevers(numbers []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	var total int
	for i, val := range numbers {
		if i == 2 {
			break
		}

		total += val
	}

	return total
}

func FindMaxSumSort(numbers []int) int {
	sort.Ints(numbers)

	var total int
	length := len(numbers)
	for i := 0; i < length; i++ {
		if i == 2 {
			break
		}

		total += numbers[length-(i+1)]
	}

	return total
}

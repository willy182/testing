package common

func FindSumNumber(findSum int, data []int) bool {
	tmp := make(map[int]int)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data)-i; j++ {
			val := data[i] + data[j]
			tmp[val] = val
		}
	}

	if _, ok := tmp[findSum]; !ok {
		return false
	}

	return true
}

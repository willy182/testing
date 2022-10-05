package common

func MinMax(data []int) (min, max int) {
	var sum int

	totalData := len(data)
	for i := 0; i < totalData-1; i++ {
		for j := 0; j < (totalData - i - 1); j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	for i := 0; i < totalData; i++ {
		sum += data[i]
	}

	min = sum - data[totalData-1]
	max = sum
	return min, max
}

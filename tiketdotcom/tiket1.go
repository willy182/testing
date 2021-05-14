package tiketdotcom

// Fractions func
func Fractions(x []int, y []int) int {
	if len(x) != len(y) {
		return 0
	}

	var total int

	pot := make(map[float64]int)
	length := len(x)

	for i := 0; i < length; i++ {
		div := float64(x[i]) / float64(y[i])
		if i == 0 {
			pot[div] = 1
			continue
		}

		if pot[div] > 0 {
			pot[div]++
		} else {
			pot[div] = 1
		}
	}

	i := 0
	for _, num := range pot {
		if i == 0 {
			total = num
			i++
			continue
		}

		if total < num {
			total = num
		}
	}

	return total
}

package tiketdotcom

// TiketTest1 func
func TiketTest1(X []int, Y []int) int {
	if len(X) != len(Y) {
		return 0
	}

	var total int

	pot := make(map[float64]int)
	length := len(X)

	for i := 0; i < length; i++ {
		div := float64(X[i]) / float64(Y[i])
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

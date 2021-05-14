package tiketdotcom

// Solution func
func Solution(N []int) int {
	totalN := len(N)
	if totalN > 3 && totalN < 201 {
		var res, j int
		var trees string

		for i := 0; i < totalN; i++ {
			if i == totalN-1 {
				break
			}

			if i == 0 {
				if (N[0] - N[j+1]) <= 0 {
					trees = "short"
				}

				if (N[0] - N[j+1]) > 0 {
					trees = "height"
				}

				j++
				res++
				continue
			}

			if (N[j] - N[i+1]) <= 0 {
				if trees != "short" {
					trees = "short"
					j = i + 1
					res++
				}

				continue
			}

			if (N[j] - N[i+1]) > 0 {
				if trees != "height" {
					trees = "height"
					j = i + 1
					res++
				}

				continue
			}
		}

		if (totalN - 1) == res {
			res = 0
		} else if res == 1 {
			res = -1
		}

		return res
	}

	return -1
}

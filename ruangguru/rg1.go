package ruangguru

import "math"

// Pangkat function per pangkatan
func Pangkat(val, n int) (result float64) {
	nAbs := int(math.Abs(float64(n)))

	var i int
	if n <= 0 {
		for i <= nAbs {
			if i == 0 {
				result = float64(val)
			}

			result /= float64(val)

			i++
		}
	} else {
		for i = 1; i <= nAbs; i++ {
			if i == 1 {
				result = 1
			}

			result *= float64(val)
		}
	}

	return result
}

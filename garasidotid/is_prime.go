package garasidotid

import (
	"math"
	"math/big"
)

// IsPrimeByFunc
// if n = prime return 1
// if n = not prime return 2
func IsPrimeByFunc(n int64) int32 {
	if big.NewInt(n).ProbablyPrime(0) {
		return 1
	}

	return 2
}

// IsPrimeBySemiNative
// if n = prime return 1
// if n = not prime return 2
func IsPrimeBySemiNative(n int64) int32 {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if int(n)%i == 0 {
			return 2
		}
	}

	if n == 1 {
		return 2
	}

	return 1
}

// IsPrimeByNative
// if n = prime return 1
// if n = not prime return 2
func IsPrimeByNative(n int64) int32 {
	for i := 2; i <= int(float64(n)/2); i++ {
		if int(n)%i == 0 {
			return 2
		}
	}

	if n == 1 {
		return 2
	}

	return 1
}

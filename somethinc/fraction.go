package somethinc

import "encoding/json"

type FractionModel struct {
	Fraction int `json:"fraction"`
	Qty      int `json:"qty"`
}

func Fractions(n int, data []int) []byte {
	res := make([]FractionModel, 0)

	for _, dt := range data {
		if dt > n {
			continue
		}

		div := n / dt
		n -= dt * div

		res = append(res, FractionModel{
			Fraction: dt,
			Qty:      div,
		})
	}

	b, _ := json.Marshal(res)

	return b
}

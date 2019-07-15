package neural

import "math/rand"

func RandomInit() WeightInit {
	return func(l, i, j int) float64 {
		return rand.Float64()
	}
}

package queen

import "math/rand"

func RandomQueen(p Population) Queens {
	q := Queens{}
	value := make([]int, p.Dim)
	for i := range value {
		value[i] = rand.Intn(p.Dim)
	}
	return q.New(value)
}

package queen

import (
	"math/rand"
	"xdean/genetic/model"
)

type Queen []int

func (q Queen) Value(i int) interface{} {
	return q[i]
}

func (q Queen) Copy() model.Single {
	result := make(Queen, len(q))
	copy(result, q)
	return result
}

func (q Queen) FindRings(o Queen) [][]int {
	ring := make([]int, len(q))
	id := 0
	for i, _ := range q {
		if ring[i] == 0 {
			id++
		}
		for {
			if ring[i] == 0 {
				ring[i] = id
				i = o.IndexOf(q[i])
			} else {
				break
			}
		}
	}
	result := make([][]int, id)
	for i := 0; i < id; i++ {
		result[i] = make([]int, 0)
	}
	for i, v := range ring {
		result[v-1] = append(result[v-1], i)
	}
	return result
}

func (q Queen) IndexOf(pos int) int {
	for i, v := range q {
		if v == pos {
			return i
		}
	}
	return -1
}

func (q Queen) RandomSwap() {
	a := rand.Intn(len(q))
	b := rand.Intn(len(q) - 1)
	if b >= a {
		b++
	}
	q[a], q[b] = q[b], q[a]
}

package sutil

import "math/rand"

func FindRings(q, o []int) [][]int {
	ring := make([]int, len(q))
	id := 0
	for i, _ := range q {
		if ring[i] == 0 {
			id++
		}
		for {
			if ring[i] == 0 {
				ring[i] = id
				i = IndexOf(o, q[i])
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

func IndexOf(q []int, pos int) int {
	for i, v := range q {
		if v == pos {
			return i
		}
	}
	return -1
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualFloat(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func RandomSwap(q []int) {
	a := rand.Intn(len(q))
	b := rand.Intn(len(q) - 1)
	if b >= a {
		b++
	}
	q[a], q[b] = q[b], q[a]
}

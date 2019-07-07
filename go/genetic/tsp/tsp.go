package tsp

import (
	"math"
	"xdean/genetic/genetic"
	"xdean/genetic/sutil"
)

type (
	Point struct {
		X float64
		Y float64
	}

	Map []Point

	TSP struct {
		Values []int
		Map    *Map
	}
)

func (p Point) Distance(o Point) float64 {
	x := p.X - o.X
	y := p.Y - o.Y
	return math.Sqrt(x*x + y*y)
}

func (q TSP) Value(i int) Point {
	m := *q.Map
	return m[q.Values[i]]
}

func (q TSP) Copy() genetic.Single {
	result := make([]int, len(q.Values))
	copy(result, q.Values)
	q.Values = result
	return q
}

func (q TSP) Equal(o genetic.Single) bool {
	switch t := o.(type) {
	case TSP:
		return q.Map == t.Map && sutil.Equal(q.Values, t.Values)
	default:
		return false
	}
}

func (q TSP) FindRings(o TSP) [][]int {
	return sutil.FindRings(q.Values, o.Values)
}

func (q TSP) IndexOf(pos int) int {
	return sutil.IndexOf(q.Values, pos)
}

func (q TSP) RandomSwap() {
	sutil.RandomSwap(q.Values)
}

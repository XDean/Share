package tsp

import (
	"math"
	"xdean/genetic/model"
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

func (q TSP) Value(i int) Point {
	m := *q.Map
	return m[q.Values[i]]
}

func (q TSP) Copy() model.Single {
	result := make([]int, len(q.Values))
	copy(result, q.Values)
	q.Values = result
	return q
}

func (p Point) Distance(o Point) float64 {
	x := p.X - o.X
	y := p.Y - o.Y
	return math.Sqrt(x*x + y*y)
}

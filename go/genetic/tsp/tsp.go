package tsp

import (
	"math"
	"xdean/genetic/model"
	"xdean/genetic/util"
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

func (q TSP) Copy() model.Single {
	result := make([]int, len(q.Values))
	copy(result, q.Values)
	q.Values = result
	return q
}

func (q TSP) FindRings(o TSP) [][]int {
	return util.FindRings(q.Values, o.Values)
}

func (q TSP) IndexOf(pos int) int {
	return util.IndexOf(q.Values, pos)
}

func (q TSP) RandomSwap() {
	util.RandomSwap(q.Values)
}

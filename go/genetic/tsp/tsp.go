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

func (p Point) Normalize(src0, src1, dst0, dst1 Point) Point {
	ratio := math.Max((src1.X-src0.X)/(dst1.X-dst0.X), (src1.Y-src0.Y)/(dst1.Y-dst0.Y))
	return Point{
		X: (p.X-src0.X)/ratio + dst0.X,
		Y: (p.Y-src0.Y)/ratio + dst0.Y,
	}
}

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
	sutil.RandomSwap(q.Values[1:])
}

func (m Map) Bounds() (x0, y0, x1, y1 float64) {
	x0 = math.MaxInt16
	y0 = math.MaxInt16
	x1 = math.MinInt16
	y1 = math.MinInt16
	for _, p := range m {
		x0 = math.Min(x0, p.X)
		y0 = math.Min(y0, p.Y)
		x1 = math.Max(x1, p.X)
		y1 = math.Max(y1, p.Y)
	}
	return
}

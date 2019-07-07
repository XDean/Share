package tsp

import (
	"math"
	"math/rand"
	"xdean/genetic/genetic"
)

func Random(m *Map) genetic.RandomFunc {
	return func(p genetic.Population) genetic.Single {
		result := TSP{
			Map:    m,
			Values: make([]int, p.Dim),
		}
		for i := range result.Values {
			result.Values[i] = i
		}
		s := result.Values[1:]
		rand.Shuffle(len(s), func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		return result
	}
}

func Crossover(p genetic.Population, ai int, bi int) (genetic.Single, genetic.Single) {
	a := p.Value[ai].(TSP)
	b := p.Value[bi].(TSP)

	return crossover1(p, a, b), crossover1(p, a, b)
}

func crossover1(p genetic.Population, a, b TSP) TSP {
	r1 := a.Copy().(TSP)
	for _, r := range a.FindRings(b) {
		if rand.Float64() > 0.5 {
			for _, i := range r {
				r1.Values[i] = a.Values[i]
			}
		} else {
			for _, i := range r {
				r1.Values[i] = b.Values[i]
			}
		}
	}
	return r1
}

func Variant(p genetic.Population, tsp genetic.Single) genetic.Single {
	new := tsp.Copy().(TSP)
	for count := p.VariantFactor / rand.Float64(); count > 0; count-- {
		new.RandomSwap()
	}
	return new
}

func ScorePow(n float64) genetic.ScoreFunc {
	return func(p genetic.Population, i int) (float64s []float64, f float64) {
		tsp := p.Value[i].(TSP)

		sum := 0.0
		length := len(tsp.Values)
		score := make([]float64, length)

		for i := 1; i < length; i++ {
			last := tsp.Value(i - 1)
			point := tsp.Value(i)
			var next Point
			if i == length-1 {
				next = tsp.Value(0)
			} else {
				next = tsp.Value(i + 1)
			}
			distance1 := math.Pow(point.Distance(last), n)
			distance2 := math.Pow(point.Distance(next), n)
			score[i] = -distance1 - distance2
			sum += -distance1 - distance2
		}

		return score, sum
	}
}

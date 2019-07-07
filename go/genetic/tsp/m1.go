package tsp

import (
	"math"
	"math/rand"
	"xdean/genetic/model"
)

func Random(m *Map) model.RandomFunc {
	return func(p model.Population) model.Single {
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

func Crossover(p model.Population, ai int, bi int) (model.Single, model.Single) {
	a := p.Value[ai].(TSP)
	b := p.Value[bi].(TSP)

	return nil, nil
}

func crossover1(p model.Population, a, b TSP) TSP {

}

func Variant(p model.Population, tsp model.Single) model.Single {
	return nil
}

func ScorePow(n float64) model.ScoreFunc {
	return func(p model.Population, i int) (float64s []float64, f float64) {
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
			score[i] = distance1 + distance2
			sum += distance1 + distance2
		}

		return score, sum
	}
}

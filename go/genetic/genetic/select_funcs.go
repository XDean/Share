package genetic

import "math/rand"

func ScoreRollSelect(p Population) int {
	r := rand.Float64() * p.TotalScore
	for q, s := range p.SingleScore {
		r -= s
		if r*p.TotalScore < 0 {
			return q
		}
	}
	panic("never happen")
}

func ScoreRollSelectTop(ratio float64) func(p Population) int {
	return func(p Population) int {
		r := rand.Float64() * ratio * p.TotalScore
		for q, s := range p.SingleScore {
			r -= s
			if r < 0 {
				return q
			}
		}
		panic("never happen")
	}
}

func ScoreOrderSelectTop(origin, less float64) func(p Population) int {
	return func(p Population) int {
		ratio := origin
		for q, _ := range p.SingleScore {
			if rand.Float64() < ratio {
				return q
			}
			ratio *= less
		}
		return 0
	}
}

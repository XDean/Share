package queen

import (
	"math"
	"math/rand"
	"xdean/genetic/model"
)

func ScoreNormalizeLossReciprocal(factor float64) func(model.ScoreFunc) model.ScoreFunc {
	return func(origin model.ScoreFunc) model.ScoreFunc {
		return func(p model.Population, i int) ([]float64, float64) {
			score, sum := origin(p, i)
			sum = factor / (1 + factor - sum)
			for i, s := range score {
				score[i] = factor / (1 + factor - s)
			}
			return score, sum
		}
	}
}

func ScorePower(n float64) model.ScoreFunc {
	return func(p model.Population, qi int) (float64s []float64, f float64) {
		q := p.Value[qi]
		score := make([]float64, len(q))
		sum := 0.0
		for c1, r1 := range q {
			for c2, r2 := range q {
				if c1 <= c2 || r1 == r2 || Abs(c1-c2) == Abs(r1-r2) {
					continue
				}
				score[c1]++
				score[c2]++
				sum++
			}
		}
		dim := float64(p.Dim)
		t := dim * (dim - 1) / 2
		sum = math.Pow(sum, n) / math.Pow(t, n)

		for i, s := range score {
			score[i] = math.Pow(s, n) / math.Pow(float64(p.Dim-1), n)
		}

		return score, sum
	}
}

func CrossoverEachGene(p model.Population, ai int, bi int) (model.Single, model.Single) {
	a := p.Value[ai]
	b := p.Value[bi]
	c1 := make([]int, p.Dim)
	c2 := make([]int, p.Dim)
	for i := 0; i < p.Dim; i++ {
		as := p.SingleGeneScore[ai][i]
		bs := p.SingleGeneScore[bi][i]
		if as/(as+bs) > rand.Float64() {
			c1[i] = a[i]
		} else {
			c1[i] = b[i]
		}
	}
	for i := 0; i < p.Dim; i++ {
		as := p.SingleGeneScore[ai][i]
		bs := p.SingleGeneScore[bi][i]
		if as/(as+bs) > rand.Float64() {
			c2[i] = a[i]
		} else {
			c2[i] = b[i]
		}
	}
	return model.NewSingle(c1), model.NewSingle(c2)
}

func CrossoverLR(p model.Population, ai int, bi int) (model.Single, model.Single) {
	a := p.Value[ai]
	b := p.Value[bi]
	c1 := make(model.Single, p.Dim)
	c2 := make(model.Single, p.Dim)

	mid := rand.Intn(p.Dim)

	copy(c1[0:mid], a[0:mid])
	copy(c1[mid:], b[mid:])
	copy(c2[0:mid], b[0:mid])
	copy(c2[mid:], a[mid:])

	return c1, c2
}

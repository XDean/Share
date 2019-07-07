package queen

import (
	"math"
	"math/rand"
	"xdean/genetic/genetic"
)

func ScoreNormalizeLossReciprocal(factor float64) func(genetic.ScoreFunc) genetic.ScoreFunc {
	return func(origin genetic.ScoreFunc) genetic.ScoreFunc {
		return func(p genetic.Population, i int) ([]float64, float64) {
			score, sum := origin(p, i)
			sum = factor / (1 + factor - sum)
			for i, s := range score {
				score[i] = factor / (1 + factor - s)
			}
			return score, sum
		}
	}
}

func ScorePower(n float64) genetic.ScoreFunc {
	return func(p genetic.Population, qi int) (float64s []float64, f float64) {
		q := p.Value[qi].(Queen)
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

func CrossoverEachGene(p genetic.Population, ai int, bi int) (genetic.Single, genetic.Single) {
	a := p.Value[ai].(Queen)
	b := p.Value[bi].(Queen)
	c1 := make(Queen, p.Dim)
	c2 := make(Queen, p.Dim)
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
	return c1, c2
}

func CrossoverLR(p genetic.Population, ai int, bi int) (genetic.Single, genetic.Single) {
	a := p.Value[ai].(Queen)
	b := p.Value[bi].(Queen)
	c1 := make(Queen, p.Dim)
	c2 := make(Queen, p.Dim)

	mid := rand.Intn(p.Dim)

	copy(c1[0:mid], a[0:mid])
	copy(c1[mid:], b[mid:])
	copy(c2[0:mid], b[0:mid])
	copy(c2[mid:], a[mid:])

	return c1, c2
}

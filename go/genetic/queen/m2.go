package queen

import (
	"math/rand"
	"xdean/genetic/model"
)

func ScoreSquare(p model.Population, qi int) ([]float64, float64) {
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
	sum = (sum * sum) / (t * t)
	return score, sum
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

package queen

import (
	"math/rand"
	"xdean/genetic/model"
)

func Random(p model.Population) model.Single {
	value := make([]int, p.Dim)
	for i := range value {
		value[i] = rand.Intn(p.Dim)
	}
	return model.NewSingle(value)
}

func Score(p model.Population, qi int) ([]float64, float64) {
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
	sum = sum / (dim * (dim - 1) / 2)
	return score, sum
}

func Crossover(p model.Population, ai int, bi int) (model.Single, model.Single) {
	a := p.Value[ai]
	b := p.Value[bi]
	c1 := make([]int, p.Dim)
	c2 := make([]int, p.Dim)
	for i := 0; i < p.Dim; i++ {
		if p.SingleGeneScore[ai][i] > p.SingleGeneScore[bi][i] {
			c1[i] = a[i]
			c2[i] = b[i]
		} else {
			c1[i] = b[i]
			c2[i] = a[i]
		}
	}
	return model.NewSingle(c1), model.NewSingle(c2)
}

func Variant(p model.Population, q model.Single) model.Single {
	per := p.VariantFactor / float64(p.Dim)
	new := q.Copy()
	for i := 0; i < p.Dim; i++ {
		r := rand.Float64()
		if r < per {
			new[i] = (new[i] + int((float64(int(1/per)%2)-0.5)*2)*(int(per/r)+1)) % p.Dim
			if new[i] < 0 {
				new[i] += p.Dim
			}
		}
	}
	return model.NewSingle(new)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package queen

import (
	"math/rand"
	"xdean/genetic/model"
)

func Random(p model.Population) model.Single {
	value := make(Queen, p.Dim)
	for i := range value {
		value[i] = i
	}
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i]
	})
	return value
}

func Score(p model.Population, qi int) ([]float64, float64) {
	q := p.Value[qi].(Queen)
	score := make([]float64, p.Dim)
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
	a := p.Value[ai].(Queen)
	b := p.Value[bi].(Queen)
	c1 := make(Queen, p.Dim)
	c2 := make(Queen, p.Dim)
	for i := 0; i < p.Dim; i++ {
		if p.SingleGeneScore[ai][i] > p.SingleGeneScore[bi][i] {
			c1[i] = a[i]
			c2[i] = b[i]
		} else {
			c1[i] = b[i]
			c2[i] = a[i]
		}
	}
	return c1, c2
}

func Variant(p model.Population, q model.Single) model.Single {
	per := p.VariantFactor / float64(p.Dim)
	new := q.Copy().(Queen)
	for i := 0; i < p.Dim; i++ {
		r := rand.Float64()
		if r < per {
			new[i] = (new[i] + int((float64(int(1/per)%2)-0.5)*2)*(int(per/r)+1)) % p.Dim
			if new[i] < 0 {
				new[i] += p.Dim
			}
		}
	}
	return new
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

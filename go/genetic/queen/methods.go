package queen

import "math/rand"

func RandomQueen(p Population) Queens {
	value := make([]int, p.Dim)
	for i := range value {
		value[i] = rand.Intn(p.Dim)
	}
	return Queens{}.New(value)
}

func ScoreQueen(p Population, q Queens) ([]float64, float64) {
	score := make([]float64, len(q.Value))
	sum := 0.0
	for c1, r1 := range q.Value {
		for c2, r2 := range q.Value {
			if c1 <= c2 || r1 == r2 || Abs(c1-c2) == Abs(r1-r2) {
				continue
			}
			score[c1]++
			score[c2]++
			sum++
		}
	}
	return score, sum
}

func CrossoverQueen(p Population, a Queens, b Queens) (Queens, Queens) {
	c1 := make([]int, p.Dim)
	c2 := make([]int, p.Dim)
	for i := 0; i < p.Dim; i++ {
		if p.SingleGeneScore[a][i] > p.SingleGeneScore[b][i] {
			c1[i] = a.Value[i]
			c2[i] = b.Value[i]
		} else {
			c1[i] = b.Value[i]
			c2[i] = a.Value[i]
		}
	}
	return Queens{}.New(c1), Queens{}.New(c2)
}

func VariantQueen(p Population, q Queens) Queens {
	per := p.VariantFactor / float64(p.Dim)
	new := q.Copy().Value
	for i := 0; i < p.Dim; i++ {
		r := rand.Float64()
		if r < per {
			new[i] = (new[i] + int((float64(int(1/per)%2)-0.5)*2)*(int(per/r)+1)) % p.Dim
			if new[i] < 0 {
				new[i] += p.Dim
			}
		}
	}
	return Queens{}.New(new)
}

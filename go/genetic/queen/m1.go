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

func CrossoverRing(p model.Population, ai int, bi int) (model.Single, model.Single) {
	a := p.Value[ai].(Queen)
	b := p.Value[bi].(Queen)

	return crossover1(p, a, b), crossover1(p, a, b)
}

func crossover1(p model.Population, a Queen, b Queen) Queen {
	r1 := make(Queen, p.Dim)
	for _, r := range a.FindRings(b) {
		if rand.Float64() > 0.5 {
			for _, i := range r {
				r1[i] = a[i]
			}
		} else {
			for _, i := range r {
				r1[i] = b[i]
			}
		}
	}
	return r1
}

func Variant(p model.Population, q model.Single) model.Single {
	new := q.Copy().(Queen)
	for count := p.VariantFactor / rand.Float64(); count > 0; count-- {
		new.RandomSwap()
	}
	return new
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

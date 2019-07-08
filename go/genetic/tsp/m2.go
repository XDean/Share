package tsp

import (
	"math/rand"
	"xdean/genetic/genetic"
)

func CrossoverRange(p genetic.Population, ai int, bi int) (genetic.Single, genetic.Single) {
	a := p.Value[ai].(TSP)
	b := p.Value[bi].(TSP)

	return crossoverRange(p, a, b), crossoverRange(p, a, b)
}

func crossoverRange(p genetic.Population, a, b TSP) TSP {
	r := a.Copy().(TSP)

	from := rand.Intn(p.Dim - 2)
	to := from + 2 + rand.Intn(p.Dim-from-2)

	for i := from; i < to; i++ {
		r.Values[i] = b.Values[i]
	}

	for ai, ri := 0, 0; ai < p.Dim; ai++ {
		bi := b.IndexOf(a.Values[ai])
		if ri == from {
			ri = to
		}
		if bi >= from && bi < to {
			continue
		} else {
			r.Values[ri] = a.Values[ai]
			ri++
		}
	}
	return r
}

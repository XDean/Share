package tsp

import (
	"math"
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

func CrossoverNearestPow(n float64) genetic.CrossoverFunc {
	return func(p genetic.Population, aIndex, bIndex int) (single genetic.Single, single2 genetic.Single) {
		crossoverNearest := func() TSP {
			a := p.Value[aIndex].(TSP)
			b := p.Value[bIndex].(TSP)
			r := a.Copy().(TSP)

			if a.Equal(b) {
				return r
			}

			use := make(map[int]bool)

			findFirst := func(resultIndex int, tsp TSP) (value int, score float64) {
				last := tsp.Value(r.Values[resultIndex-1])
				ai := (tsp.IndexOf(r.Values[resultIndex-1]) + 1) % p.Dim
				for {
					if ai == 0 {
						ai++
					}
					av := tsp.Values[ai]
					au := use[av]
					current := tsp.Value(ai)
					distance := last.Distance(current)
					if au {
						ai = (ai + 1) % p.Dim
					} else {
						return av, math.Pow(distance, n)
					}
				}
			}

			for i := 1; i < p.Dim; i++ {
				av, as := findFirst(i, a)
				bv, bs := findFirst(i, b)
				rd := rand.Float64() * (as + bs)
				if rd < as {
					use[av] = true
					r.Values[i] = av
				} else {
					use[bv] = true
					r.Values[i] = bv
				}
			}
			return r
		}
		return crossoverNearest(), crossoverNearest()
	}
}

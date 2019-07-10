package tsp

import (
	"github.com/xdean/share/go/genetic/genetic"
	"math"
	"math/rand"
)

func VariantRevertSwap(p genetic.Population, tsp genetic.Single) genetic.Single {
	new := tsp.Copy().(TSP)
	for count := p.VariantFactor / rand.Float64(); count > 1; count-- {
		from := rand.Intn(p.Dim-1) + 1
		to := rand.Intn(p.Dim-1) + 1
		new.Revert(from, to)
	}
	return new
}

func CrossoverNearestRevert(pow float64) genetic.CrossoverFunc {
	return func(p genetic.Population, aIndex, bIndex int) (single genetic.Single, single2 genetic.Single) {
		crossoverNearest := func() TSP {
			a := p.Value[aIndex].Copy().(TSP)
			b := p.Value[bIndex].Copy().(TSP)

			if a.Equal(b) {
				return a
			}
			aUsed := make(map[int]bool)

			findBNearest := func(aIndex int) int {
				av := a.Values[aIndex]
				bi := b.IndexOf(av)
				for i := 1; ; i++ {
					bn := (bi + i) % p.Dim
					if !aUsed[b.Values[bn]] {
						return bn
					}
					bl := (bi + p.Dim - i) % p.Dim
					if !aUsed[b.Values[bl]] {
						return bl
					}
				}
			}

			for i := 1; i < p.Dim; i++ {
				bIndex := findBNearest(i - 1)
				aToIndex := a.IndexOf(b.Values[bIndex])
				aToIndex2 := (aToIndex + 1) % p.Dim

				distanceOld := math.Pow(a.Value(i-1).Distance(a.Value(i))+a.Value(aToIndex).Distance(a.Value(aToIndex2)), pow)
				distanceNew := math.Pow(a.Value(i-1).Distance(a.Value(aToIndex))+a.Value(i).Distance(a.Value(aToIndex2)), pow)
				rd := rand.Float64() * (distanceOld + distanceNew)
				if rd < distanceNew {
					a.Revert(i, aToIndex)
				}
				aUsed[a.Values[i]] = true
			}
			return a
		}
		return crossoverNearest(), crossoverNearest()
	}
}

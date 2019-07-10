package function

import (
	"math"
	"math/rand"
	"xdean/share/genetic/genetic"
)

func Random(f Function) genetic.RandomFunc {
	return func(p genetic.Population) genetic.Single {
		result := Input{
			Func:  f,
			Value: make([]float64, p.Dim),
		}
		for i, _ := range result.Value {
			result.Value[i] = rand.Float64()
		}
		return result
	}
}

func ScorePow(baseInput, targetOutput []float64, inputFactor, pow float64) genetic.ScoreFunc {
	return func(p genetic.Population, i int) ([]float64, float64) {
		input := p.Value[i].(Input)
		output := input.Output()

		score := make([]float64, p.Dim)
		outputScore := 0.0
		sum := 0.0

		for i, v := range targetOutput {
			outputScore += math.Pow(1-math.Abs(output[i]-v), pow)
		}
		outputScore /= float64(len(targetOutput))

		for i, v := range baseInput {
			score[i] = (outputScore + math.Pow(1-math.Abs(input.Value[i]-v)*inputFactor, pow)) / 2
			sum += score[i]
		}
		sum /= float64(p.Dim)
		return score, sum
	}
}

func Crossover() genetic.CrossoverFunc {
	return func(p genetic.Population, ai, bi int) (genetic.Single, genetic.Single) {
		crossover := func() Input {
			a := p.Value[ai].(Input)
			b := p.Value[bi].(Input)
			r := a.Copy().(Input)
			for i, av := range a.Value {
				bv := b.Value[i]
				as := p.SingleGeneScore[ai][i]
				bs := p.SingleGeneScore[bi][i]
				if rand.Float64()*(as+bs) < as {
					r.Value[i] = av
				} else {
					r.Value[i] = bv
				}
			}
			return r
		}
		return crossover(), crossover()
	}
}

func Variant(step float64) genetic.VariantFunc {
	return func(p genetic.Population, s genetic.Single) genetic.Single {
		r := s.Copy().(Input)
		for i := range r.Value {
			if rand.Float64() < p.VariantFactor/float64(p.Dim) {
				r.Value[i] = math.Max(0, math.Min(1, r.Value[i]+(2*rand.Float64()-1)*step))
			}
		}
		return r
	}
}

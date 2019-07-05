package model

import (
	"math/rand"
)

type (
	RandomFunc    func(Population) Single
	CrossoverFunc func(Population, int, int) (Single, Single)
	VariantFunc   func(Population, Single) Single
	ScoreFunc     func(Population, int) ([]float64, float64)

	Population struct {
		Gen             int
		Size            int
		Dim             int
		CrossoverFactor float64
		VariantFactor   float64
		Target          float64
		MaxGen          int

		RandomFunc    RandomFunc
		CrossoverFunc CrossoverFunc
		VariantFunc   VariantFunc
		ScoreFunc     ScoreFunc

		Value           []Single
		TotalScore      float64
		SingleGeneScore [][]float64
		SingleScore     []float64
	}
)

func (p Population) Random() Population {
	value := make([]Single, p.Size)
	for i := range value {
		value[i] = p.RandomFunc(p)
	}
	return p.Brother(value).Score()
}

func (p Population) CanContinue() bool {
	return p.Gen < p.MaxGen
}

func (p Population) IsDone() (ok bool, maxScore float64, maxValue Single) {
	for i, score := range p.SingleScore {
		if score >= p.Target {
			return true, score, p.Value[i]
		}
		if score > maxScore {
			maxValue = p.Value[i]
			maxScore = score
		}
	}
	return false, maxScore, maxValue
}

func (p Population) Brother(value []Single) Population {
	p.Value = value
	p.SingleScore = make([]float64, p.Size)
	p.SingleGeneScore = make([][]float64, p.Size)
	p.TotalScore = 0
	return p.Score()
}

func (p Population) Child(value []Single) Population {
	bro := p.Brother(value)
	bro.Gen++
	return bro
}

func (p Population) Score() Population {
	sum := 0.0
	for i, _ := range p.Value {
		per, total := p.ScoreFunc(p, i)
		p.SingleGeneScore[i] = per
		p.SingleScore[i] = total
		sum += total
	}
	p.TotalScore = sum
	return p
}

func (p Population) NextGen() Population {
	output := make(chan Single, 10)
	for i := 0; i < p.Size; i += 2 {
		go p.next2(output)
	}

	new := make([]Single, p.Size)
	for i := 0; i < p.Size; i++ {
		new[i] = <-output
	}
	close(output)
	return p.Child(new)
}

func (p Population) next2(output chan<- Single) {
	q1 := p.next()
	q2 := p.next()

	var r1, r2 Single

	if rand.Float64() < p.CrossoverFactor {
		r1, r2 = p.CrossoverFunc(p, q1, q2)
	} else {
		r1, r2 = p.Value[q1].Copy(), p.Value[q2].Copy()
	}
	r1 = p.VariantFunc(p, r1)
	r2 = p.VariantFunc(p, r2)
	output <- r1
	output <- r2
}

func (p Population) next() int {
	r := rand.Float64() * p.TotalScore
	for q, s := range p.SingleScore {
		r -= s
		if r < 0 {
			return q
		}
	}
	panic("never happen")
}

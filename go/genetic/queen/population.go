package queen

import "math/rand"

type Population struct {
	Gen             int
	Size            int
	Dim             int
	CrossoverFactor float64
	VariantFactor   float64
	Target          int

	RandomFunc    func(Population) Queens
	CrossoverFunc func(Population, Queens, Queens) (Queens, Queens)
	VariantFunc   func(Population, Queens) Queens
	ScoreFunc     func(Population, Queens) ([]float64, float64)

	Value           []Queens
	TotalScore      float64
	SingleGeneScore map[Queens][]float64
	SingleScore     map[Queens]float64
}

func (p Population) Random() Population {
	value := make([]Queens, p.Size)
	for i := range value {
		value[i] = p.RandomFunc(p)
	}
	result := p.Brother(value)
	result.Score()
	return result
}

func (p Population) Brother(value []Queens) Population {
	p.Value = value
	p.SingleScore = map[Queens]float64{}
	p.SingleGeneScore = map[Queens][]float64{}
	p.TotalScore = 0
	p.Score()
	return p
}

func (p Population) Child(value []Queens) Population {
	bro := p.Brother(value)
	bro.Gen++
	return bro
}

func (p Population) Score() {
	sum := 0.0
	for _, q := range p.Value {
		per, total := p.ScoreFunc(p, q)
		p.SingleGeneScore[q] = per
		p.SingleScore[q] = total
		sum += total
	}
	p.TotalScore = sum
}

func (p Population) NextGen() Population {
	output := make(chan Queens, 10)
	for i := 0; i < p.Size; i += 2 {
		go p.next2(output)
	}

	new := make([]Queens, p.Size)
	for i := 0; i < p.Size; i++ {
		new[i] = <-output
	}
	close(output)
	return p.Child(new)
}

func (p Population) next2(output chan<- Queens) {
	q1 := p.next()
	q2 := p.next()

	if rand.Float64() < p.CrossoverFactor {
		q1, q2 = p.CrossoverFunc(p, q1, q2)
	}
	q1 = p.VariantFunc(p, q1)
	q2 = p.VariantFunc(p, q2)
	output <- q1
	output <- q2
}

func (p Population) next() Queens {
	r := rand.Intn(p.TotalScore)
	for _, q := range p.Value {
		r -= q.TotalScore
		if r < 0 {
			return q
		}
	}
	panic("never happen")
}

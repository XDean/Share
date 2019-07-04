package queen

import "math/rand"

type Population struct {
	Gen             int
	Size            int
	Dim             int
	CrossoverFactor float64
	VariantFactor   float64
	Target          int

	Value      []Queens
	TotalScore int
}

func (p Population) Random() Population {
	value := make([]Queens, p.Size)
	for i := range value {
		value[i] = Queens{}.Random(p.Dim)
	}
	return p.Brother(value)
}

func (p Population) Brother(value []Queens) Population {
	p.Value = value
	total := 0
	for _, q := range p.Value {
		total += q.TotalScore
	}
	p.TotalScore = total
	return p
}

func (p Population) Child(value []Queens) Population {
	bro := p.Brother(value)
	bro.Gen++
	return bro
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
		q1, q2 = q1.Crossover(q2)
	}
	q1 = q1.Variant(p.VariantFactor)
	q2 = q2.Variant(p.VariantFactor)
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

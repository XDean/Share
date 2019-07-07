package genetic

import (
	"math/rand"
	"sort"
)

type (
	RandomFunc    func(Population) Single
	CrossoverFunc func(Population, int, int) (Single, Single)
	VariantFunc   func(Population, Single) Single
	ScoreFunc     func(Population, int) ([]float64, float64)
	SelectFunc    func(Population) int
	TargetFunc    func(Population) bool
	PluginFunc    func(Population) Population
	Plugin        struct {
		Start PluginFunc
		Each  PluginFunc
		End   PluginFunc
	}

	Population struct {
		Gen             int
		Size            int
		Dim             int
		CrossoverFactor float64
		VariantFactor   float64
		MaxGen          int

		RandomFunc    RandomFunc
		CrossoverFunc CrossoverFunc
		VariantFunc   VariantFunc
		ScoreFunc     ScoreFunc
		SelectFunc    SelectFunc
		TargetFunc    TargetFunc
		Plugins       []Plugin

		// all below always sort by score descending
		Value           []Single
		TotalScore      float64
		SingleGeneScore [][]float64
		SingleScore     []float64
	}
)

func (p Population) Run() Population {
	for _, plugin := range p.Plugins {
		p = plugin.Start(p)
	}
	for _, plugin := range p.Plugins {
		p = plugin.Each(p)
	}
	for p.NeedContinue() {
		p = p.NextGen()
		for _, plugin := range p.Plugins {
			p = plugin.Each(p)
		}
	}
	for _, plugin := range p.Plugins {
		p = plugin.End(p)
	}
	return p
}

func (p Population) Random() Population {
	value := make([]Single, p.Size)
	for i := range value {
		value[i] = p.RandomFunc(p)
	}
	return p.Brother(value).Score()
}

func (p Population) NeedContinue() bool {
	return p.Gen < p.MaxGen && !p.MatchTarget()
}

func (p Population) MatchTarget() bool {
	return p.TargetFunc(p)
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

	sort.Sort(p)
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
	q1 := p.SelectFunc(p)
	q2 := p.SelectFunc(p)

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

func (p Population) Len() int {
	return p.Size
}

func (p Population) Less(i, j int) bool {
	return p.SingleScore[i] > p.SingleScore[j]
}

func (p Population) Swap(i, j int) {
	p.Value[i], p.Value[j] = p.Value[j], p.Value[i]
	p.SingleScore[i], p.SingleScore[j] = p.SingleScore[j], p.SingleScore[i]
	p.SingleGeneScore[i], p.SingleGeneScore[j] = p.SingleGeneScore[j], p.SingleGeneScore[i]
}

/*****KPI******/
func (p Population) BestSingle() Single {
	return p.Value[0]
}

func (p Population) BestScore() float64 {
	return p.SingleScore[0]
}

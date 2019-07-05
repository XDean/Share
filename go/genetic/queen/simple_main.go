package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/model"
)

func SimpleMain() {
	rand.Seed(time.Now().Unix())
	dim := 10

	population := model.Population{
		Size:            100,
		Dim:             dim,
		CrossoverFactor: 0.8,
		VariantFactor:   0.1,
		Target:          0.999,
		MaxGen:          5000,

		RandomFunc:    Random,
		CrossoverFunc: Crossover,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePower(1),
	}.Random()

	model.CalcAndPlotBox(population, "points.svg")
}

package queen

import "xdean/genetic/model"

func SimpleMain() {
	dim := 8

	population := model.Population{
		Size:            100,
		Dim:             dim,
		CrossoverFactor: 0.9,
		VariantFactor:   0.1,
		Target:          0.999,
		MaxGen:          5000,

		RandomFunc:    Random,
		CrossoverFunc: Crossover,
		VariantFunc:   Variant,
		ScoreFunc:     ScoreSquare,
	}.Random()

	model.CalcAndPlotBox(population, "points.svg")
}

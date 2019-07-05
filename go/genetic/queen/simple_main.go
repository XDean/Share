package queen

import "xdean/genetic/model"

func SimpleMain() {
	dim := 20

	population := model.Population{
		Size:            100,
		Dim:             dim,
		CrossoverFactor: 0.8,
		VariantFactor:   0.2,
		Target:          0.999,
		MaxGen:          5000,

		RandomFunc:    Random,
		CrossoverFunc: CrossoverLR,
		VariantFunc:   Variant,
		ScoreFunc:     Score,
	}.Random()

	model.CalcAndPlotBox(population, "points.svg")
}

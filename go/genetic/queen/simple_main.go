package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/model"
)

func GeneticMain() {
	rand.Seed(time.Now().Unix())
	dim := 200

	population := model.Population{
		Size:            500,
		Dim:             dim,
		CrossoverFactor: 1,
		VariantFactor:   0.2,
		Target:          1,
		MaxGen:          5000,

		RandomFunc:    Random,
		CrossoverFunc: CrossoverRing,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePower(10),
		SelectFunc:    model.ScoreOrderSelectTop(0.05, 0.9),
	}.Random()

	model.CalcAndPlotBox(population, "points.svg")
}

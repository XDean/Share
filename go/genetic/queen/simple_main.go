package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/model"
)

func SimpleMain() {
	rand.Seed(time.Now().Unix())
	dim := 20

	population := model.Population{
		Size:            200,
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

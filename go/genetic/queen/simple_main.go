package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/model"
)

func SimpleMain() {
	rand.Seed(time.Now().Unix())
	dim := 15

	population := model.Population{
		Size:            500,
		Dim:             dim,
		CrossoverFactor: 1,
		VariantFactor:   0.2,
		Target:          0.999,
		MaxGen:          5000,

		RandomFunc:    Random,
		CrossoverFunc: CrossoverRing,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePower(5),
	}.Random()

	model.CalcAndPlotBox(population, "points.svg")
}

package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/genetic"
)

func GeneticMain() {
	rand.Seed(time.Now().Unix())
	dim := 200

	population := genetic.Population{
		Size:            500,
		Dim:             dim,
		CrossoverFactor: 1,
		VariantFactor:   0.05,
		MaxGen:          5000,

		TargetFunc:    genetic.TargetScore(1),
		RandomFunc:    Random,
		CrossoverFunc: CrossoverRing,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePower(1),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.05, 0.9),
	}.Random()

	genetic.CalcAndPlotBox(population, "points.svg")
}

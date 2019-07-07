package tsp

import (
	"math/rand"
	"time"
	"xdean/genetic/genetic"
)

func Main() {
	rand.Seed(time.Now().Unix())

	tspMap := Map{
		{X: 116.46, Y: 39.92},
		{X: 117.20, Y: 39.13},
		{X: 121.48, Y: 31.22},
		{X: 106.54, Y: 29.59},
	}

	population := genetic.Population{
		Size:            500,
		Dim:             len(tspMap),
		CrossoverFactor: 1,
		VariantFactor:   0.05,
		MaxGen:          5000,

		TargetFunc:    genetic.TargetScore(1),
		RandomFunc:    Random(&tspMap),
		CrossoverFunc: Crossover,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePow(2),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.05, 0.9),
	}.Random()

	genetic.CalcAndPlotBox(population, "output/tsp.svg")
}

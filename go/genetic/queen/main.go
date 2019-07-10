package queen

import (
	"math/rand"
	"time"
	"xdean/genetic/genetic"
	"xdean/genetic/genetic/plugin"
)

func GeneticMain() {
	rand.Seed(time.Now().Unix())
	dim := 200

	genetic.Population{
		Size:            100,
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

		Plugins: []genetic.Plugin{
			plugin.Print(),
			plugin.Timing(),
			plugin.BoxPlot("Queen by GA", "output/queen.svg"),
			plugin.ImagePerGenBest("output/queen/", ToImage, 100000),
		},
	}.Random().Run()
}

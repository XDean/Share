package queen

import (
	"github.com/xdean/share/go/genetic/genetic"
	"github.com/xdean/share/go/genetic/genetic/plugin"
	"math/rand"
	"time"
)

func GeneticMain() {
	rand.Seed(time.Now().Unix())
	dim := 20

	genetic.Population{
		Size:            100,
		Dim:             dim,
		CrossoverFactor: 0.8,
		VariantFactor:   0.2,
		MaxGen:          5000,

		TargetFunc:    genetic.TargetScore(1),
		RandomFunc:    Random,
		CrossoverFunc: CrossoverRing,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePower(1),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.1, 0.9),

		Plugins: []genetic.Plugin{
			plugin.Print(false),
			plugin.Timing(),
			plugin.BoxPlot("Queen by GA", "output/queen.svg"),
			plugin.ImagePerGenBest("output/queen/", ToImage, 1),
		},
	}.Random().Run()
}

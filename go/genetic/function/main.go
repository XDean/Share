package function

import (
	"fmt"
	"math/rand"
	"time"
	"xdean/genetic/genetic"
	"xdean/genetic/genetic/plugin"
)

func Main() {
	rand.Seed(time.Now().Unix())

	var function Function = func(input []float64) []float64 {
		return []float64{
			(input[0] + input[1] - input[2]),
		}
	}

	baseInput := []float64{0.5, 0.5, 0.5}
	targetOutput := []float64{0.6}

	result := genetic.Population{
		Size:            50,
		Dim:             len(baseInput),
		CrossoverFactor: 0.9,
		VariantFactor:   0.2,
		MaxGen:          2000,

		TargetFunc:    genetic.TargetStableScore(100).Or(genetic.TargetScore(1)),
		RandomFunc:    Random(function),
		CrossoverFunc: Crossover(),
		VariantFunc:   Variant(0.1),
		ScoreFunc:     ScorePow(baseInput, targetOutput, 0.3, 2),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.1, 0.8),

		Plugins: []genetic.Plugin{
			plugin.Print(),
			plugin.BoxPlot("Function Solution by GA", "output/function.svg"),
		},
	}.Random().Run()
	answer := result.BestSingle().(Input)
	fmt.Println()
	fmt.Println("Target Output", targetOutput)
	fmt.Println("Function Output", function(answer.Value))
	fmt.Println()
	fmt.Println()
}

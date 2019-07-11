package function

import (
	"fmt"
	"github.com/xdean/share/go/genetic/genetic"
	"github.com/xdean/share/go/genetic/genetic/plugin"
	"math/rand"
	"time"
)

func Main() {
	rand.Seed(time.Now().Unix())

	var function Function = func(input []float64) []float64 {
		return []float64{
			input[0] + input[1] - input[2],
			input[0] * input[1] * input[2],
		}
	}

	baseInput := []float64{0.5, 0.5, 0.5}
	targetOutput := []float64{0.51, 0.13}

	result := genetic.Population{
		Size:            10,
		Dim:             len(baseInput),
		CrossoverFactor: 0.9,
		VariantFactor:   0.2,
		MaxGen:          2000,

		TargetFunc:    genetic.TargetStableScore(10).Or(genetic.TargetScore(1)),
		RandomFunc:    Random(function),
		CrossoverFunc: Crossover(),
		VariantFunc:   Variant(0.1),
		ScoreFunc:     ScorePow(baseInput, targetOutput, 0.3, 2),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.1, 0.8),

		Plugins: []genetic.Plugin{
			plugin.PrintEachBest(),
			plugin.Timing(),
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

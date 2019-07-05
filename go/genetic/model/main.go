package model

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)
import "gonum.org/v1/plot"

func CalcAndPlotBox(population Population, outputFile string) {

	plot, err := plot.New()
	if err != nil {
		panic(err)
	}
	plot.Title.Text = "Queen by GA"
	plot.X.Label.Text = "Gen"
	plot.Y.Label.Text = "Score"

	grid := plotter.NewGrid()
	grid.Horizontal.Width = 0
	grid.Vertical.Width = 1

	plot.Add(grid)

	result := Single{}
	score := 0.0

	for population.CanContinue() {
		scores := make(plotter.Values, population.Size)
		for i, s := range population.SingleScore {
			scores[i] = s
		}
		box, err := plotter.NewBoxPlot(10, float64(population.Gen), scores)
		if err != nil {
			panic(err)
		}
		plot.Add(box)

		if ok, maxScore, max := population.IsDone(); ok {
			result = max
			score = maxScore
			break
		} else {
			fmt.Printf("Gen %d, total score %.2f, best score %.2f, value %v \n", population.Gen, population.TotalScore, maxScore, max)
		}
		population = population.NextGen()
	}

	fmt.Println("Total Gen", population.Gen)
	fmt.Println("Score", score)
	fmt.Println("Answer", result)

	if err := plot.Save(vg.Length(population.Gen*15), vg.Length(600), outputFile); err != nil {
		panic(err)
	}
}

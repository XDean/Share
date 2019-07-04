package queen

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)
import "gonum.org/v1/plot"

func Main() {

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

	population := Population{
		Size:            1000,
		Dim:             15,
		CrossoverFactor: 0.8,
		VariantFactor:   0.2,
		Target:          105,
	}.Random()
	result := Queens{}
outside:
	for {
		scores := make(plotter.Values, population.Size)
		for i, q := range population.Value {
			scores[i] = float64(q.TotalScore)
		}
		box, err := plotter.NewBoxPlot(10, float64(population.Gen), scores)
		if err != nil {
			panic(err)
		}
		plot.Add(box)

		max := Queens{}
		for _, q := range population.Value {
			if q.TotalScore >= population.Target {
				result = q
				break outside
			}
			if q.TotalScore > max.TotalScore {
				max = q
			}
		}
		fmt.Printf("Gen %d, best score %d, value %v \n", population.Gen, max.TotalScore, max.Value)
		population = population.NextGen()
	}
	fmt.Println("Total Gen", population.Gen)
	fmt.Println("Answer", result.Value)

	if err := plot.Save(vg.Length(population.Gen*15), vg.Length(population.Target*5), "points.svg"); err != nil {
		panic(err)
	}
}

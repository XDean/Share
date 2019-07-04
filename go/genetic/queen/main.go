package queen

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
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
	points := make(plotter.XYs, 0)

	population := Population{
		Size:            100,
		Dim:             8,
		CrossoverFactor: 0.9,
		VariantFactor:   0.3,
		Target:          28,
	}.Random()
	result := Queens{}
outside:
	for {
		for _, q := range population.Value {
			points = append(points, plotter.XY{X: float64(population.Gen), Y: float64(q.TotalScore)})
		}
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

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	plot.Add(scatter)

	if err := plot.Save(1024, 1024, "points.svg"); err != nil {
		panic(err)
	}
}

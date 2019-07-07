package model

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)
import "gonum.org/v1/plot"

func CalcAndPlotBox(p Population, outputFile string) {

	plot, err := plot.New()
	if err != nil {
		panic(err)
	}
	plot.Title.Text = "Queen by GA"
	plot.X.Label.Text = "Gen"
	plot.Y.Label.Text = "Score"

	totalScores := plotter.XYs{}

	for p.NeedContinue() {
		scores := make(plotter.Values, p.Size)
		for i, s := range p.SingleScore {
			scores[i] = s
		}
		box, err := plotter.NewBoxPlot(10, float64(p.Gen), scores)
		if err != nil {
			panic(err)
		}
		plot.Add(box)
		totalScores = append(totalScores, plotter.XY{X: float64(p.Gen), Y: p.TotalScore / float64(p.Size)})

		fmt.Printf("Gen %d, total score %.2f, best score %.2f, value %v \n", p.Gen, p.TotalScore, p.BestScore(), p.BestSingle())

		p = p.NextGen()
	}

	fmt.Println("Total Gen", p.Gen)
	fmt.Println("Find Target", p.MatchTarget())
	fmt.Println("Best Score", p.BestScore())
	fmt.Println("Best Answer", p.BestSingle())

	err = plotutil.AddLinePoints(plot, totalScores)
	if err != nil {
		panic(err)
	}
	if err := plot.Save(vg.Length(p.Gen*20+100), vg.Length(600), outputFile); err != nil {
		panic(err)
	}
}

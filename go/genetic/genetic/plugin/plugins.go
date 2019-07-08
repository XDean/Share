package plugin

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"path/filepath"
	"xdean/genetic/genetic"
)

func Print() genetic.Plugin {
	return genetic.Plugin{
		Start: func(p genetic.Population) genetic.Population {
			fmt.Println("Start GA")
			return p
		},
		Each: func(p genetic.Population) genetic.Population {
			fmt.Printf("Gen %d, total score %.2f, best score %.4f, value %v \n", p.Gen, p.TotalScore, p.BestScore(), p.BestSingle())
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			fmt.Println("Total Gen", p.Gen)
			fmt.Println("Find Target", p.MatchTarget())
			fmt.Println("Best Score", p.BestScore())
			fmt.Println("Best Answer", p.BestSingle())
			return p
		},
	}
}

func BoxPlot(title, path string) genetic.Plugin {
	plot, err := plot.New()
	if err != nil {
		panic(err)
	}
	plot.Title.Text = title
	plot.X.Label.Text = "Gen"
	plot.Y.Label.Text = "Score"

	totalScores := plotter.XYs{}

	return genetic.Plugin{
		Start: genetic.EMPTY_PLUGIN_FUNC,
		Each: func(p genetic.Population) genetic.Population {
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
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			err = plotutil.AddLinePoints(plot, totalScores)
			if err != nil {
				panic(err)
			}
			_ = os.MkdirAll(filepath.Dir(path), os.ModeType)
			if err := plot.Save(vg.Length(p.Gen*20+100), vg.Length(600), path); err != nil {
				panic(err)
			}
			return p
		},
	}
}

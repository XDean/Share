package plugin

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func Timing() genetic.Plugin {
	start := time.Now()
	return genetic.Plugin{
		Start: func(p genetic.Population) genetic.Population {
			start = time.Now()
			return p
		},
		Each: genetic.EMPTY_PLUGIN_FUNC,
		End: func(p genetic.Population) genetic.Population {
			time := time.Since(start)
			fmt.Printf("Total Take %s\n", time)
			return p
		},
	}
}

func BoxPlot(title, path string) genetic.Plugin {
	pt, err := plot.New()
	if err != nil {
		panic(err)
	}
	pt.Title.Text = title
	pt.X.Label.Text = "Gen"
	pt.Y.Label.Text = "Score"
	ticks := make([]plot.Tick, 0)
	pt.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		return ticks
	})

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
			pt.Add(box)
			ticks = append(ticks, plot.Tick{Value: float64(p.Gen), Label: strconv.Itoa(p.Gen)})
			totalScores = append(totalScores, plotter.XY{X: float64(p.Gen), Y: p.TotalScore / float64(p.Size)})
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			err = plotutil.AddLinePoints(pt, totalScores)
			if err != nil {
				panic(err)
			}
			_ = os.MkdirAll(filepath.Dir(path), os.ModeType)
			if err := pt.Save(vg.Length(p.Gen*20+100), vg.Length(600), path); err != nil {
				panic(err)
			}
			return p
		},
	}
}

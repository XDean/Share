package neural

import (
	"fmt"
	"golang.org/x/image/colornames"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"math"
	"math/rand"
	"os"
	"testing"
)

func TestModel(t *testing.T) {
	inputSize := 2
	outputSize := 1
	model := Model{
		Config: ModelConfig{
			LayerCount:   4,
			NodeCount:    []int{inputSize, 6, 3, outputSize},
			LearningRate: 0.1,
			Activation:   ReLU,
		},
	}
	model.Init()

	for i := 0; i < 10000; i++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		model.Feed([]float64{x, y}, []float64{x*x + y*y - 0.5 + x*math.Sin(x+y)})
		fmt.Printf("Train Loss %.2f, %.2f, %.5f\n", x, y, model.Value.Error)
	}
	fmt.Println("0.1,0.2", "->", model.Predict([]float64{0.1, 0.2}))
	fmt.Println("-0.1,0.2", "->", model.Predict([]float64{-0.1, 0.2}))

	_ = os.MkdirAll("output", os.ModeType)
	for i := 0; i < outputSize; i++ {
		pt, err := plot.New()
		panicErr(err)
		positive := make(plotter.XYs, 0)
		negative := make(plotter.XYs, 0)
		for x := -1.0; x <= 1.0; x += 0.05 {
			for y := -1.0; y <= 1.0; y += 0.05 {
				if model.Predict([]float64{x, y})[i] > 0 {
					positive = append(positive, plotter.XY{x, y})
				} else {
					negative = append(negative, plotter.XY{x, y})
				}
			}
		}
		ps, err := plotter.NewScatter(positive)
		panicErr(err)
		ps.GlyphStyle.Color = colornames.Red
		ns, err := plotter.NewScatter(negative)
		panicErr(err)
		ns.Shape = draw.CrossGlyph{}
		ns.GlyphStyle.Color = colornames.Blue
		pt.Add(ps, ns)
		if err := pt.Save(vg.Length(600), vg.Length(600), fmt.Sprintf("output/ann_%d.svg", i)); err != nil {
			panic(err)
		}
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

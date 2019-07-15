package neural

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestModel(t *testing.T) {
	model := Model{
		Config: ModelConfig{
			LayerCount:   4,
			NodeCount:    []int{2, 6, 3, 1},
			LearningRate: 0.1,
			Activation:   ReLU,
		},
	}
	model.Init()

	for i := 0; i < 100; i++ {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5
		o := x * y * 100
		model.Feed([]float64{x, y}, []float64{o})
		fmt.Printf("Train Loss %.5f\n", model.Value.Error)

	}
	fmt.Println("0.1,0.2", "->", model.Predict([]float64{0.1, 0.2})[0])
	fmt.Println("-0.1,0.2", "->", model.Predict([]float64{-0.1, 0.2})[0])
}

package queen

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"testing"
	"xdean/genetic/genetic"
)

func TestQueen_ToImage(t *testing.T) {
	var queen Queen = []int{2, 1, 3, 5, 6, 0, 7, 4}
	img := ToImage(genetic.Population{Gen: 123, Value: []genetic.Single{queen}, SingleScore: []float64{0.99}}, 0)
	draw2dimg.SaveToPngFile("../output/queen_test.png", img)
}

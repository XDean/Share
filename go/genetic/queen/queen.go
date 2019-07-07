package queen

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image"
	"image/color"
	"xdean/genetic/genetic"
	"xdean/genetic/sutil"
)

type Queen []int

func (q Queen) Copy() genetic.Single {
	result := make(Queen, len(q))
	copy(result, q)
	return result
}

func (q Queen) Equal(o genetic.Single) bool {
	switch t := o.(type) {
	case Queen:
		return sutil.Equal(q, t)
	default:
		return false
	}
}

func (q Queen) FindRings(o Queen) [][]int {
	return sutil.FindRings(q, o)
}

func (q Queen) IndexOf(pos int) int {
	return sutil.IndexOf(q, pos)
}

func (q Queen) RandomSwap() {
	sutil.RandomSwap(q)
}

func (q Queen) ToImage() image.Image {
	len := len(q)
	width := len*30 + 60
	height := len*30 + 90

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(img)

	for row, column := range q {
		for i := 0; i < len; i++ {
			if i+row%2 == 0 {
				gc.SetFillColor(color.RGBA{R: 0xff, G: 0x66, B: 0x66, A: 0xaa})
			} else {
				gc.SetFillColor(color.RGBA{R: 0xff, G: 0x66, B: 0x66, A: 0x88})
			}
			x := float64(i*30 + 30)
			y := float64(row*30 + 60)
			gc.MoveTo(x, y)
			gc.LineTo(x+30, y)
			gc.LineTo(x+30, y+30)
			gc.LineTo(x, y+30)
			gc.Close()
			gc.Fill()
		}
		x := float64(column*30 + 45)
		y := float64(row*30 + 75)
		draw2dkit.Circle(gc, x, y, 10)
		gc.Stroke()
	}
	return img
}

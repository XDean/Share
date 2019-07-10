package queen

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/xdean/share/go/genetic/genetic"
	"image"
	"image/color"
	"image/draw"
)

func ToImage(p genetic.Population, index int) image.Image {
	q := p.Value[index].(Queen)
	len := len(q)
	width := len*30 + 60
	height := len*30 + 90

	gc := gg.NewContext(width, height)
	gc.DrawRectangle(0, 0, float64(width), float64(height))
	gc.SetColor(color.White)
	gc.Fill()
	gc.SetColor(color.Black)
	gc.DrawStringAnchored(fmt.Sprintf("Queen Gen %d, Score %.4f", p.Gen, p.SingleScore[index]), float64(width/2), 30, 0.5, 0.5)

	for row, column := range q {
		for i := 0; i < len; i++ {
			if (i+row)%2 == 0 {
				gc.SetColor(color.Gray{Y: 188})
			} else {
				gc.SetColor(color.Gray{Y: 222})
			}
			x := float64(i*30 + 30)
			y := float64(row*30 + 60)
			gc.DrawRectangle(x, y, 30, 30)
			gc.Fill()
		}
		x := float64(column*30 + 45)
		y := float64(row*30 + 75)
		gc.SetColor(color.Black)
		gc.DrawCircle(x, y, 10)
		gc.Fill()
	}

	img := gc.Image()
	grayImg := image.NewGray(img.Bounds())
	draw.Draw(grayImg, grayImg.Rect, img, image.ZP, draw.Src)
	return grayImg
}

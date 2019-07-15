package main

import (
	"github.com/xdean/share/go/neural/neural"
	"image/png"
	"os"
)

func DigitReadImage(imgFile string) []float64 {
	imgReader, err := os.Open(imgFile)
	neural.PanicErr(err)
	img, err := png.Decode(imgReader)
	neural.PanicErr(err)
	input := make([]float64, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			color := img.At(i, j)
			r, g, b, _ := color.RGBA()
			if r+g+b > 255 {
				input[i*28+j] = 0
			} else {
				input[i*28+j] = 1
			}
		}
	}
	return input
}

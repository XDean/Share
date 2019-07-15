package main

import (
	"fmt"
	"github.com/xdean/goex/xgo"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"testing"
)

func TestGenGif(t *testing.T) {
	file, err := os.Create("tsp.gif")
	xgo.MustNoError(err)
	defer file.Close()

	imgs := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	for i := 0; i <= 32; i++ {
		imgFile, err := os.Open(fmt.Sprintf("%04d.png", i))
		if err != nil {
			fmt.Println("No Image", i)
			continue
		}
		img, err := png.Decode(imgFile)
		if err != nil {
			fmt.Println("Can't read image", i)
			continue
		}
		palImg := image.NewPaletted(img.Bounds(), color.Palette{color.Black, color.White})
		draw.Draw(palImg, palImg.Bounds(), img, image.ZP, draw.Src)
		imgs = append(imgs, palImg)
		delays = append(delays, 100)
	}

	err = gif.EncodeAll(file, &gif.GIF{
		Image: imgs,
		Delay: delays,
	})
	xgo.MustNoError(err)
}

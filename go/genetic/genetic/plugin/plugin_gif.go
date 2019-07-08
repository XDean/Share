package plugin

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"
	"xdean/genetic/genetic"
)

type ImageFunc func(population genetic.Population, index int) image.Image

func Gif(path string, delay float64, imageFunc ImageFunc) genetic.Plugin {
	var images []*image.Paletted
	var delays []int

	return genetic.Plugin{
		Start: genetic.EMPTY_PLUGIN_FUNC,
		Each: func(p genetic.Population) genetic.Population {
			img := imageFunc(p, 0)
			palette := image.NewPaletted(img.Bounds(), color.Palette{})
			draw.Draw(palette, palette.Rect, img, image.ZP, draw.Src)
			images = append(images, palette)
			delays = append(delays, 0)
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			_ = os.MkdirAll(filepath.Dir(path), os.ModeType)
			f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
			defer f.Close()
			err := gif.EncodeAll(f, &gif.GIF{
				Image: images,
				Delay: delays,
			})
			if err != nil {
				panic(err)
			}
			return p
		},
	}
}

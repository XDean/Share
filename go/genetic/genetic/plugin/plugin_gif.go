package plugin

import (
	"image"
	"image/gif"
	"os"
	"xdean/genetic/genetic"
)

type Imagable interface {
	ToImage() image.Image
}

func Gif(path string, delay float64) genetic.Plugin {
	var images []image.Image
	var delays []int

	return genetic.Plugin{
		Start: genetic.EMPTY_PLUGIN_FUNC,
		Each: func(p genetic.Population) genetic.Population {
			images = append(images, p.BestSingle().(Imagable).ToImage())
			delays = append(delays, 0)
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
			defer f.Close()
			gif.EncodeAll(f, &gif.GIF{
				Image: images,
				Delay: delays,
			})
			return p
		},
	}
}

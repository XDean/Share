package plugin

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"path/filepath"
	"sync"
	"xdean/genetic/genetic"
)

type ImageFunc func(population genetic.Population, index int) image.Image

func ImageEachBest(folder string, imageFunc ImageFunc) genetic.Plugin {
	return ImagePerGenBest(folder, imageFunc, 1)
}

func ImagePerGenBest(folder string, imageFunc ImageFunc, gen int) genetic.Plugin {
	err := os.MkdirAll(folder, os.ModeType)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	var last genetic.Single
	return genetic.Plugin{
		Start: genetic.EMPTY_PLUGIN_FUNC,
		Each: func(p genetic.Population) genetic.Population {
			if p.Gen%gen != 0 || p.Value[0].Equal(last) {
				return p
			} else {
				last = p.Value[0]
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				img := imageFunc(p, 0)
				path := filepath.Join(folder, fmt.Sprintf("%04d.png", p.Gen))
				if file, err := os.Create(path); err != nil {
					panic(err)
				} else {
					defer file.Close()
					err := png.Encode(file, img)
					if err != nil {
						panic(err)
					}
				}
			}()
			return p
		},
		End: func(p genetic.Population) genetic.Population {
			wg.Wait()
			return p
		},
	}
}

// Not available
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

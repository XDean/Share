package tsp

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
	"xdean/genetic/genetic"
)

func Random(m *Map) genetic.RandomFunc {
	return func(p genetic.Population) genetic.Single {
		result := TSP{
			Map:    m,
			Values: make([]int, p.Dim),
		}
		for i := range result.Values {
			result.Values[i] = i
		}
		s := result.Values[1:]
		rand.Shuffle(len(s), func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		return result
	}
}

func CrossoverRing(p genetic.Population, ai int, bi int) (genetic.Single, genetic.Single) {
	a := p.Value[ai].(TSP)
	b := p.Value[bi].(TSP)

	return crossover1(p, a, b), crossover1(p, a, b)
}

func crossover1(p genetic.Population, a, b TSP) TSP {
	r1 := a.Copy().(TSP)
	for _, r := range a.FindRings(b) {
		if rand.Float64() > 0.5 {
			for _, i := range r {
				r1.Values[i] = a.Values[i]
			}
		} else {
			for _, i := range r {
				r1.Values[i] = b.Values[i]
			}
		}
	}
	return r1
}

func VariantSwap(p genetic.Population, s genetic.Single) genetic.Single {
	tsp := s.(TSP)
	new := tsp.Copy().(TSP)
	if rand.Float64() < p.VariantFactor {
		cut := rand.Intn(p.Dim-2) + 2
		copy(new.Values[1:p.Dim-cut+1], tsp.Values[cut:p.Dim])
		copy(new.Values[p.Dim-cut+1:p.Dim], tsp.Values[1:cut])
	}
	return new
}

func Variant(p genetic.Population, tsp genetic.Single) genetic.Single {
	new := tsp.Copy().(TSP)
	for count := p.VariantFactor / rand.Float64(); count > 2; count-- {
		new.RandomSwap()
		new.RandomSwap()
	}
	return new
}

func ScoreDistancePow(n float64) genetic.ScoreFunc {
	return func(p genetic.Population, i int) (float64s []float64, f float64) {
		tsp := p.Value[i].(TSP)

		sum := 0.0
		length := len(tsp.Values)
		score := make([]float64, length)
		score[0] = 1

		for i := 1; i < length; i++ {
			last := tsp.Value(i - 1)
			point := tsp.Value(i)
			var next Point
			if i == length-1 {
				next = tsp.Value(0)
			} else {
				next = tsp.Value(i + 1)
			}
			d1 := point.Distance(last)
			d2 := point.Distance(next)
			distance1 := math.Pow(d1, n)
			distance2 := math.Pow(d2, n)
			score[i] = distance1 + distance2
			sum += d1
			if i == length-1 {
				sum += d2
			}
		}
		sum = math.Pow(sum, n)

		return score, sum
	}
}

func ToImage(p genetic.Population, index int) image.Image {
	tsp := p.Value[index].(TSP)
	x0, y0, x1, y1 := tsp.Map.Bounds()

	src0 := Point{x0, y0}
	src1 := Point{x1, y1}
	dst0 := Point{30, 60}
	dst1 := Point{970, 770}

	width := 1000
	height := 800

	gc := gg.NewContext(int(width), int(height))
	gc.DrawRectangle(0, 0, float64(width), float64(height))
	gc.SetColor(color.White)
	gc.Fill()

	gc.DrawRectangle(10, 40, float64(width-20), float64(height-60))
	gc.SetColor(color.Black)
	gc.Stroke()

	gc.SetColor(color.Black)
	gc.DrawStringAnchored(fmt.Sprintf("TSP Gen %d, Score %.4f", p.Gen, p.SingleScore[index]), float64(width/2), 15, 0.5, 0.5)

	drawPointAndLine := func(current int, last Point, point Point) {
		last = last.Normalize(src0, src1, dst0, dst1)
		point = point.Normalize(src0, src1, dst0, dst1)
		gc.DrawCircle(point.X, 820-point.Y, 5)
		gc.Fill()
		gc.DrawStringAnchored(fmt.Sprintf("%d", current), point.X, 800-point.Y, 0.5, 0.5)
		gc.DrawLine(last.X, 820-last.Y, point.X, 820-point.Y)
		gc.Stroke()
	}

	for i := 1; i < len(tsp.Values); i++ {
		last := tsp.Value(i - 1)
		point := tsp.Value(i)
		drawPointAndLine(tsp.Values[i], last, point)
	}
	drawPointAndLine(tsp.Values[0], tsp.Value(len(tsp.Values)-1), tsp.Value(0))

	img := gc.Image()
	grayImg := image.NewGray(img.Bounds())
	draw.Draw(grayImg, grayImg.Rect, img, image.ZP, draw.Src)
	return grayImg
}

package tsp

import (
	"math/rand"
	"time"
	"xdean/genetic/genetic"
	"xdean/genetic/genetic/plugin"
)

func Main() {
	rand.Seed(time.Now().Unix())

	tspMap := Map{
		{X: 116.46, Y: 39.92},
		{X: 117.20, Y: 39.13},
		{X: 121.48, Y: 31.22},
		{X: 106.54, Y: 29.59},
		{X: 091.11, Y: 29.97},
		{X: 087.68, Y: 43.77},
		{X: 106.27, Y: 38.47},
		{X: 111.65, Y: 40.82},
		{X: 108.33, Y: 22.84},
		{X: 126.63, Y: 45.75},
		{X: 125.35, Y: 43.88},
		{X: 123.38, Y: 41.80},
		{X: 114.48, Y: 38.03},
		{X: 112.53, Y: 37.87},
		{X: 101.74, Y: 36.56},
		{X: 117.00, Y: 36.65},
		{X: 113.60, Y: 34.76},
		{X: 117.27, Y: 31.86},
		{X: 120.19, Y: 30.26},
		{X: 119.3, Y: 26.08},
		{X: 115.89, Y: 28.68},
		{X: 113.00, Y: 28.21},
		{X: 114.31, Y: 30.52},
		{X: 113.23, Y: 23.16},
		{X: 121.5, Y: 25.05},
		{X: 110.35, Y: 20.02},
		{X: 103.73, Y: 36.03},
		{X: 108.95, Y: 34.27},
		{X: 104.06, Y: 30.67},
		{X: 106.71, Y: 26.57},
		{X: 102.73, Y: 25.04},
		{X: 114.1, Y: 22.2},
		{X: 113.33, Y: 22.13},
	}

	genetic.Population{
		Size:            500,
		Dim:             len(tspMap),
		CrossoverFactor: 1,
		VariantFactor:   0.05,
		MaxGen:          200,

		TargetFunc:    genetic.TargetScore(1),
		RandomFunc:    Random(&tspMap),
		CrossoverFunc: Crossover,
		VariantFunc:   Variant,
		ScoreFunc:     ScorePow(2),
		SelectFunc:    genetic.ScoreOrderSelectTop(0.05, 0.9),

		Plugins: []genetic.Plugin{
			plugin.Print(),
			plugin.BoxPlot("TSP by GA", "output/tsp.svg"),
			plugin.ImageEachBest("output/tsp", ToImage),
		},
	}.Random().Run()
}

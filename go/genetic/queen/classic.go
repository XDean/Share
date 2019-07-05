package queen

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

var results = make([][]Point, 0)

func ClassicMain() {
	Solve(14)
}

func Solve(n int) {
	defer func() {
		_ = recover()
		fmt.Print("Results:\n")
		for _, result := range results {
			fmt.Println(result)
		}
		fmt.Printf("There were %d results\n", len(results))
	}()
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
}
func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results = append(results, c)
			panic("not panic")
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)
				}
			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	//fmt.Print(a, b)
	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))
	//fmt.Print(answer)
	return answer
}

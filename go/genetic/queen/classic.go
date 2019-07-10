package queen

import (
	"fmt"
	"math"
	"sync"
)

type Point struct {
	x int
	y int
}

var results = make([][]Point, 1)

func ClassicMain() {
	Solve(14)
}

func Solve(n int) {
	defer func() {
		fmt.Print("Results:\n")
		for _, result := range results {
			fmt.Println(result)
		}
	}()
	wg := sync.WaitGroup{}
	for col := 0; col < n; col++ {
		wg.Add(1)
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		go func() {
			Recurse(start, current, n)
			wg.Done()
		}()
	}
	wg.Wait()
}
func Recurse(point Point, current []Point, n int) {
	if results[0] != nil {
		return
	}
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results[0] = c
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

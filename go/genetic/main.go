package main

import (
	"fmt"
	"time"
	"xdean/genetic/queen"
)

func main() {
	start := time.Now()

	//queen.ClassicMain()
	queen.GeneticMain()

	elapsed := time.Since(start)

	fmt.Printf("Take %s\n", elapsed)
}

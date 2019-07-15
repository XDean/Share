package queen

import (
	"fmt"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	start := time.Now()

	//ClassicMain()
	GeneticMain()

	elapsed := time.Since(start)

	fmt.Printf("Take %s\n", elapsed)
}

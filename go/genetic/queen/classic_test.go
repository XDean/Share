package queen

import (
	"fmt"
	"testing"
	"time"
)

func Test_Classic(t *testing.T) {
	start := time.Now()

	ClassicMain()

	elapsed := time.Since(start)

	fmt.Printf("Take %s\n", elapsed)
}

package queen

import (
	"github.com/xdean/share/go/genetic/genetic"
	"github.com/xdean/share/go/genetic/sutil"
)

type Queen []int

func (q Queen) Copy() genetic.Single {
	result := make(Queen, len(q))
	copy(result, q)
	return result
}

func (q Queen) Equal(o genetic.Single) bool {
	switch t := o.(type) {
	case Queen:
		return sutil.Equal(q, t)
	default:
		return false
	}
}

func (q Queen) FindRings(o Queen) [][]int {
	return sutil.FindRings(q, o)
}

func (q Queen) IndexOf(pos int) int {
	return sutil.IndexOf(q, pos)
}

func (q Queen) RandomSwap() {
	sutil.RandomSwap(q)
}

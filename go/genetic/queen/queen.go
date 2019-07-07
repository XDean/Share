package queen

import (
	"xdean/genetic/model"
	"xdean/genetic/util"
)

type Queen []int

func (q Queen) Copy() model.Single {
	result := make(Queen, len(q))
	copy(result, q)
	return result
}

func (q Queen) FindRings(o Queen) [][]int {
	return util.FindRings(q, o)
}

func (q Queen) IndexOf(pos int) int {
	return util.IndexOf(q, pos)
}

func (q Queen) RandomSwap() {
	util.RandomSwap(q)
}

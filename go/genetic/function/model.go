package function

import (
	"xdean/genetic/genetic"
	"xdean/genetic/sutil"
)

type (
	Function *func([]float64) []float64
	Input    struct {
		Value []float64
		Func  Function
	}
)

func (i Input) Copy() genetic.Single {
	return i
}

func (i Input) Equal(o genetic.Single) bool {
	switch t := o.(type) {
	case Input:
		return i.Func == t.Func && sutil.EqualFloat(i.Value, t.Value)
	default:
		return false
	}
}

func (i Input) Output() []float64 {
	return (*i.Func)(i.Value)
}

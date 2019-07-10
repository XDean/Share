package function

import (
	"xdean/share/genetic/genetic"
	"xdean/share/genetic/sutil"
)

type (
	Function func([]float64) []float64
	Input    struct {
		Value []float64
		Func  Function
	}
)

func (i Input) Copy() genetic.Single {
	result := make([]float64, len(i.Value))
	copy(result, i.Value)
	i.Value = result
	return i
}

func (i Input) Equal(o genetic.Single) bool {
	switch t := o.(type) {
	case Input:
		return sutil.EqualFloat(i.Value, t.Value)
	default:
		return false
	}
}

func (i Input) Output() []float64 {
	return i.Func(i.Value)
}

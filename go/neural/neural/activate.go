package neural

import "math"

type (
	Activation interface {
		Active(input float64) (output, partial float64)
	}
	ActivationFunc func(input float64) (output, partial float64)
)

var (
	Sigmoid ActivationFunc = func(input float64) (output, partial float64) {
		output = 1 / (1 + math.Exp(-input))
		return output, output * (1 - output)
	}
	ReLU ActivationFunc = func(input float64) (output, partial float64) {
		if input > 0 {
			return input, 1
		} else {
			return 0, 0
		}
	}
)

func (a ActivationFunc) Active(input float64) (output, partial float64) {
	return a(input)
}

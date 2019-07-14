package neural

import "errors"

type (
	Model struct {
		Config struct {
			LayerCount int
			NodeCount  []int
			Activation Activation
		}
		Value struct {
			Error  float64
			Node   [][]Node
			Weight [][][]Weight
			Delta  [][]float64
		}
	}

	Node struct {
		Input   float64
		Output  float64
		Partial float64
	}

	Weight struct {
		Value float64
		Nabla float64
	}
)

func (m Model) Feed(input []float64) error {
	if len(input) != m.Config.NodeCount[0] {
		return errors.New("wrong input length")
	}
	for i := range m.Value.Node[0] {
		m.Value.Node[0][i].Input = input[i]
	}
	return nil
}

func (m Model) Forward() {
	for l := 1; l < m.Config.LayerCount; l++ {
		for i, in := range m.Value.Node[l] {
			sum := 0.0
			for k, kn := range m.Value.Node[l-1] {
				sum += kn.Output * m.Value.Weight[l][k][i].Value
			}
			in.Input = sum
			in.Output, in.Partial = m.Config.Activation.Active(in.Input)
		}
	}
}

func (m Model) CalcError() {

}

func (m Model) CalcGradient() {

}

package neural

import "errors"

type (
	Model struct {
		Config struct {
			LayerCount   int
			NodeCount    []int
			Activation   Activation
			LearningRate float64
		}
		Value struct {
			Error  float64
			Node   [][]Node
			Weight [][][]Weight // first index starts from 1
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

func (m Model) Init() {

}

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
			if l == m.Config.LayerCount-1 {
				in.Output, in.Partial = in.Input, 1
			} else {
				in.Output, in.Partial = m.Config.Activation.Active(in.Input)
			}
		}
	}
}

func (m Model) CalcError(target []float64) {
	m.Value.Error = 0
	for i, v := range m.Value.Node[m.Config.LayerCount-1] {
		m.Value.Error += (v.Output - target[i]) * (v.Input - target[i]) / 2
	}
}

func (m Model) Backward(target []float64) {
	for l := m.Config.LayerCount - 1; l > 0; l-- {
		for j := 0; j < m.Config.NodeCount[l]; j++ {
			if l == m.Config.LayerCount-1 {
				m.Value.Delta[l][j] = (m.Value.Node[l][j].Output - target[j]) * m.Value.Node[l][j].Partial
			} else {
				sum := 0.0
				for k := 0; k < m.Config.NodeCount[l+1]; k++ {
					sum += m.Value.Delta[l+1][k] * m.Value.Weight[l+1][j][k].Value
				}
				m.Value.Delta[l][j] = sum * m.Value.Node[l][j].Partial
			}
			for i := 0; i < m.Config.NodeCount[l-1]; i++ {
				m.Value.Weight[l][i][j].Nabla = m.Value.Delta[l][j] * m.Value.Node[l-1][i].Output
			}
		}
	}
}

func (m Model) Learn() {
	for l := 1; l < m.Config.LayerCount; l++ {
		for i := 0; i < m.Config.NodeCount[l-1]; i++ {
			for j := 0; j < m.Config.NodeCount[l]; j++ {
				m.Value.Weight[l][i][j].Value -= m.Config.LearningRate * m.Value.Weight[l][i][j].Nabla
			}
		}
	}
}

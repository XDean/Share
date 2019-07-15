package neural

type (
	Model struct {
		Config struct {
			LayerCount   int
			NodeCount    []int
			Activation   Activation
			LearningRate float64
			WeightInit   WeightInit
		}
		Value struct {
			Error  float64
			Node   [][]Node
			Weight [][][]Weight // first index starts from 1
		}
	}

	WeightInit func(l, i, j int) float64

	Node struct {
		Input   float64
		Output  float64
		Partial float64 // d output / d input
		Delta   float64 // δ
	}

	Weight struct {
		Value float64
		Nabla float64 // δ * output[l-1][i]
	}
)

func (m Model) Init() {
	m.Value.Node = make([][]Node, m.Config.LayerCount)
	for l := range m.Value.Node {
		m.Value.Node[l] = make([]Node, m.Config.NodeCount[l])
	}
	m.Value.Weight = make([][][]Weight, m.Config.LayerCount)
	for l := 1; l < m.Config.LayerCount; l++ {
		m.Value.Weight[l] = make([][]Weight, m.Config.NodeCount[l-1])
		for i := 0; i < m.Config.NodeCount[l-1]; i++ {
			m.Value.Weight[l][i] = make([]Weight, m.Config.NodeCount[l])
		}
	}
	if m.Config.WeightInit == nil {
		m.Config.WeightInit = RandomInit()
	}
	for l := 1; l < m.Config.LayerCount; l++ {
		for i := 0; i < m.Config.NodeCount[l-1]; i++ {
			for j := 0; j < m.Config.NodeCount[l]; j++ {
				m.Value.Weight[l][i][j].Value = m.Config.WeightInit(l, i, j)
			}
		}
	}
}

func (m Model) Feed(input, output []float64) error {
	m.Input(input)
	m.Forward()
	m.CalcError(output)
	m.Backward(output)
	m.Learn()
	return nil
}

func (m Model) Input(input []float64) {
	if len(input) != m.Config.NodeCount[0] {
		panic("wrong input length")
	}
	for i := range m.Value.Node[0] {
		m.Value.Node[0][i].Input = input[i]
	}
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
				m.Value.Node[l][j].Delta = (m.Value.Node[l][j].Output - target[j]) * m.Value.Node[l][j].Partial
			} else {
				sum := 0.0
				for k := 0; k < m.Config.NodeCount[l+1]; k++ {
					sum += m.Value.Node[l+1][k].Delta * m.Value.Weight[l+1][j][k].Value
				}
				m.Value.Node[l][j].Delta = sum * m.Value.Node[l][j].Partial
			}
			for i := 0; i < m.Config.NodeCount[l-1]; i++ {
				m.Value.Weight[l][i][j].Nabla = m.Value.Node[l][j].Delta * m.Value.Node[l-1][i].Output
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

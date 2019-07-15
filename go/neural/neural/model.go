package neural

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

type (
	Model struct {
		Config ModelConfig
		Value  ModelValue
	}

	WeightInit func(l, i, j int) float64

	ModelConfig struct {
		LayerCount   int
		NodeCount    []int
		Activation   Activation
		LearningRate float64
		WeightInit   WeightInit
	}

	ModelValue struct {
		Error  float64
		Node   [][]Node
		Weight [][][]Weight // first index starts from 1, second index end to node count + 1 (last one is bias)
	}

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

func (m *Model) Save(file string) error {
	err := os.MkdirAll(filepath.Dir(file), os.ModeDir)
	if err != nil {
		return err
	}
	writer, err := os.Create(file)
	if err != nil {
		return err
	}
	defer writer.Close()
	encoder := gob.NewEncoder(writer)
	err = encoder.Encode(m.Config.LayerCount)
	if err != nil {
		return err
	}
	err = encoder.Encode(m.Config.NodeCount)
	if err != nil {
		return err
	}
	err = encoder.Encode(m.Value.Weight)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) Load(file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	defer reader.Close()
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&m.Config.LayerCount)
	if err != nil {
		return err
	}
	err = decoder.Decode(&m.Config.NodeCount)
	if err != nil {
		return err
	}
	m.Init()
	err = decoder.Decode(&m.Value.Weight)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) Init() {
	m.Value.Node = make([][]Node, m.Config.LayerCount)
	for l := range m.Value.Node {
		m.Value.Node[l] = make([]Node, m.Config.NodeCount[l])
	}
	m.Value.Weight = make([][][]Weight, m.Config.LayerCount)
	for l := 1; l < m.Config.LayerCount; l++ {
		m.Value.Weight[l] = make([][]Weight, m.Config.NodeCount[l-1]+1)
		for i := 0; i <= m.Config.NodeCount[l-1]; i++ {
			m.Value.Weight[l][i] = make([]Weight, m.Config.NodeCount[l])
		}
	}
	if m.Config.WeightInit == nil {
		m.Config.WeightInit = RandomInit()
	}
	for l := 1; l < m.Config.LayerCount; l++ {
		for j := 0; j < m.Config.NodeCount[l]; j++ {
			for i := 0; i < m.Config.NodeCount[l-1]; i++ {
				m.Value.Weight[l][i][j].Value = m.Config.WeightInit(l, i, j)
			}
			m.Value.Weight[l][m.Config.NodeCount[l-1]][j].Value = 0.1
		}
	}
}

func (m *Model) Feed(input, output []float64) {
	m.Input(input)
	m.Forward()
	m.CalcError(output)
	m.Backward(output)
	m.Learn()
}

func (m *Model) Test(input, output []float64) float64 {
	m.Input(input)
	m.Forward()
	m.CalcError(output)
	return m.Value.Error
}

func (m *Model) Predict(input []float64) []float64 {
	m.Input(input)
	m.Forward()
	result := make([]float64, m.Config.NodeCount[m.Config.LayerCount-1])
	for i, v := range m.Value.Node[m.Config.LayerCount-1] {
		result[i] = v.Output
	}
	return result
}

func (m *Model) Input(input []float64) {
	if len(input) != m.Config.NodeCount[0] {
		panic("wrong input length")
	}
	for i := range m.Value.Node[0] {
		m.Value.Node[0][i].Input = input[i]
		m.Value.Node[0][i].Output = input[i]
		m.Value.Node[0][i].Partial = 1
	}
}

func (m *Model) Forward() {
	for l := 1; l < m.Config.LayerCount; l++ {
		for i := range m.Value.Node[l] {
			in := &m.Value.Node[l][i]
			sum := 0.0
			for k, kn := range m.Value.Node[l-1] {
				sum += kn.Output * m.Value.Weight[l][k][i].Value
			}
			sum += m.Value.Weight[l][m.Config.NodeCount[l-1]][i].Value
			in.Input = sum
			if l == m.Config.LayerCount-1 {
				in.Output, in.Partial = in.Input, 1
			} else {
				in.Output, in.Partial = m.Config.Activation.Active(in.Input)
			}
		}
	}
}

func (m *Model) CalcError(target []float64) {
	m.Value.Error = 0
	for i, v := range m.Value.Node[m.Config.LayerCount-1] {
		m.Value.Error += (v.Output - target[i]) * (v.Output - target[i]) / 2
	}
}

func (m *Model) Backward(target []float64) {
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
			m.Value.Weight[l][m.Config.NodeCount[l-1]][j].Nabla = m.Value.Node[l][j].Delta
		}
	}
}

func (m *Model) Learn() {
	for l := 1; l < m.Config.LayerCount; l++ {
		for i := 0; i <= m.Config.NodeCount[l-1]; i++ {
			for j := 0; j < m.Config.NodeCount[l]; j++ {
				m.Value.Weight[l][i][j].Value -= m.Config.LearningRate * m.Value.Weight[l][i][j].Nabla
			}
		}
	}
}

func (w Weight) String() string {
	return fmt.Sprintf("%.3f", w.Value)
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

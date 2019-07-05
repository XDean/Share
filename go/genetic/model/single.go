package model

type (
	Single []int
)

func NewSingle(value []int) Single {
	q := make([]int, len(value))
	copy(q, value)
	return q
}

func (q Single) Copy() Single {
	return NewSingle(q)
}

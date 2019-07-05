package queen

type (
	Queens struct {
		Value []int
	}
)

func (q Queens) New(value []int) Queens {
	q.Value = make([]int, len(value))
	copy(q.Value, value)
	return q
}

func (q Queens) Copy() Queens {
	return Queens{}.New(q.Value)
}

func Assert(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

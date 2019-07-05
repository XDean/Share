package queen

import "xdean/genetic/model"

type Queen []int

func (q Queen) Value(i int) interface{} {
	return q[i]
}

func (q Queen) Copy() model.Single {
	result := make(Queen, len(q))
	copy(result, q)
	return result
}

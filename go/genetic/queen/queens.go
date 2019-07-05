package queen

import "math/rand"

type (
	Queens struct {
		Size       int
		Value      []int
		Score      []int
		TotalScore int
	}
)

func (q Queens) New(value []int) Queens {
	q.Size = len(value)
	q.Value = make([]int, len(value))
	copy(q.Value, value)
	score := make([]int, len(q.Value))
	sum := 0
	for c1, r1 := range q.Value {
		for c2, r2 := range q.Value {
			if c1 <= c2 || r1 == r2 || Abs(c1-c2) == Abs(r1-r2) {
				continue
			}
			score[c1]++
			score[c2]++
			sum++
		}
	}
	q.Score = score
	q.TotalScore = sum
	return q
}

func (q Queens) Crossover(o Queens) (Queens, Queens) {
	Assert(q.Size == o.Size, "Need same size")
	c1 := make([]int, q.Size)
	c2 := make([]int, q.Size)
	for i := 0; i < q.Size; i++ {
		if q.Score[i] > o.Score[i] {
			c1[i] = q.Value[i]
			c2[i] = o.Value[i]
		} else {
			c1[i] = o.Value[i]
			c2[i] = q.Value[i]
		}
	}
	return Queens{}.New(c1), Queens{}.New(c2)
}

func (q Queens) Copy() Queens {
	return Queens{}.New(q.Value)
}

func (q Queens) Variant(factor float64) Queens {
	per := factor / float64(q.Size)
	new := q.Copy().Value
	for i := 0; i < q.Size; i++ {
		r := rand.Float64()
		if r < per {
			new[i] = (new[i] + int((float64(int(1/per)%2)-0.5)*2)*(int(per/r)+1)) % q.Size
			if new[i] < 0 {
				new[i] += q.Size
			}
		}
	}
	return Queens{}.New(new)
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

package model

type (
	Single interface {
		Value(int) interface{}
		Copy() Single
	}
)

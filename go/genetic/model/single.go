package model

type (
	Single interface {
		Copy() Single
		Equal(Single) bool
	}
)

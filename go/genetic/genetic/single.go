package genetic

type (
	Single interface {
		Copy() Single
		Equal(Single) bool
	}
)

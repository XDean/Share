package genetic

func TargetScore(s float64) TargetFunc {
	return func(p Population) bool {
		return p.BestScore() >= s
	}
}

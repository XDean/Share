package genetic

func TargetScore(s float64) TargetFunc {
	return func(p Population) bool {
		return p.BestScore() >= s
	}
}

func TargetStableScore(gen int) TargetFunc {
	lastGen := -1
	lastScore := -1.0
	return func(p Population) bool {
		s := p.BestScore()
		if s > lastScore {
			lastScore = s
			lastGen = p.Gen
		} else if p.Gen-lastGen > gen {
			return true
		}
		return false
	}
}

func (t TargetFunc) Or(o TargetFunc) TargetFunc {
	return func(p Population) bool {
		return t(p) || o(p)
	}
}

func (t TargetFunc) And(o TargetFunc) TargetFunc {
	return func(p Population) bool {
		return t(p) && o(p)
	}
}

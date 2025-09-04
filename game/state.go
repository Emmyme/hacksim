package game

type State struct {
	Score int
	Level int
}

func NewState(score int, level int) *State {
	return &State{
		Score: score,
		Level: level,
	}
}
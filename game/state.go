package game

type State struct {
	Name string
	Score int
	Level int
}

func NewState(name string, score int, level int) *State {
	return &State{
		Name:  name,
		Score: score,
		Level: level,
	}
}
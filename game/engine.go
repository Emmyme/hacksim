package game

type Engine struct {
	state          *State
	puzzles        []Puzzle
	currentPuzzle  int
	missions       []Mission
	currentMission int
	commands       []Command
	currentCommand int
}

func NewEngine(state *State) *Engine {
	return &Engine{
		state: state,
	}
}
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
		state:          state,
		puzzles:        GetPuzzles(),
		missions:       GetMissions(),
		commands:       GetCommands(),
		currentPuzzle:  0,
		currentMission: 0,
		currentCommand: 0,
	}
}

func (e *Engine) CurrentPuzzle() int {
	return e.currentPuzzle
}

func (e *Engine) CurrentMission() int {
	return e.currentMission
}

func (e *Engine) CurrentCommand() int {
	return e.currentCommand
}

func (e *Engine) Puzzles() []Puzzle {
	return e.puzzles
}

func (e *Engine) Missions() []Mission {
	return e.missions
}

func (e *Engine) Commands() []Command {
	return e.commands
}

func (e *Engine) NextMission() {
	if e.currentMission < len(e.missions)-1 {
		e.currentMission++
	}
}

func (e *Engine) NextPuzzle() {
	if e.currentPuzzle < len(e.puzzles)-1 {
		e.currentPuzzle++
	}
}

func (e *Engine) State() *State {
	return e.state
}

func (e *Engine) NextIncompleteMission() *Mission {
	for i := range e.missions {
		if !e.missions[i].Completed {
			e.currentMission = i
			return &e.missions[i]
		}
	}
	return nil
}

func (e *Engine) NextIncompletePuzzle() *Puzzle {
	for i := range e.puzzles {
		if !e.puzzles[i].Completed {
			e.currentPuzzle = i
			return &e.puzzles[i]
		}
	}
	return nil
}

func (e *Engine) AllMissionsCompleted() bool {
	for _, mission := range e.missions {
		if !mission.Completed {
			return false
		}
	}
	return true
}

func (e *Engine) AllPuzzlesCompleted() bool {
	completed := 0
	for _, puzzle := range e.puzzles {
		if puzzle.Completed {
			completed++
		}
	}
	return completed == len(e.puzzles)
}

func (e *Engine) MarkCurrentPuzzleCompleted() {
	if e.currentPuzzle < len(e.puzzles) {
		e.puzzles[e.currentPuzzle].Completed = true
	}
}

func (e *Engine) MarkCurrentMissionCompleted() {
	if e.currentMission < len(e.missions) {
		e.missions[e.currentMission].Completed = true
	}
}

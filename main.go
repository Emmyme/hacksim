package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Emmyme/hacksim/game"
)

type model struct {
	engine     *game.Engine
	input      string
	mode       string
	discovered bool
	gameWon    bool
}

func (m *model) resetGame() {
	m.discovered = true
	m.engine.State().Score = 0
	m.engine.State().Level = 1
	m.engine = game.NewEngine(m.engine.State())
	m.mode = "mission"
}

func (m *model) checkWinCondition() {
	if m.engine.AllMissionsCompleted() && m.engine.AllPuzzlesCompleted() {
		m.gameWon = true
		m.mode = "won"
	} else {
		m.mode = "puzzle"
	}
}

func (m *model) checkWinConditionAfterPuzzle() {
	if m.engine.AllMissionsCompleted() && m.engine.AllPuzzlesCompleted() {
		m.gameWon = true
		m.mode = "won"
	} else {
		m.mode = "mission"
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.HideCursor,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.mode == "mission" && m.engine.AllMissionsCompleted() && !m.engine.AllPuzzlesCompleted() {
				m.mode = "puzzle"
				m.input = ""
				return m, nil
			}

			if m.mode == "puzzle" && m.engine.AllPuzzlesCompleted() && !m.engine.AllMissionsCompleted() {
				m.mode = "mission"
				m.input = ""
				return m, nil
			}

			if m.engine.AllMissionsCompleted() && m.engine.AllPuzzlesCompleted() {
				m.gameWon = true
				m.mode = "won"
				m.input = ""
				return m, nil
			}

			if m.mode == "mission" {
				currentMission := m.engine.NextIncompleteMission()
				if currentMission != nil {
					commands := m.engine.Commands()
					validCommand := false
					for _, cmd := range commands {
						if cmd.Name == m.input {
							cmd.Execute(m.engine.State())
							validCommand = true
							break
						}
					}
					if !validCommand {
						m.resetGame()
					} else {
						correctSequence := currentMission.Progress(m.input)
						if !correctSequence {
							m.resetGame()
						} else {
							m.discovered = false
							if currentMission.Completed {
								m.engine.State().Level++
								m.checkWinCondition()
							}
						}
					}
				}
				m.input = ""
			} else if m.mode == "puzzle" {
				currentPuzzle := m.engine.NextIncompletePuzzle()
				if currentPuzzle != nil {
					if m.input == currentPuzzle.Solution {
						m.discovered = false
						m.engine.State().Score += 20
						m.engine.State().Level++
						m.engine.MarkCurrentPuzzleCompleted()
						m.checkWinConditionAfterPuzzle()
					} else {
						m.resetGame()
					}
				}
				m.input = ""
			}
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5"))
	scoreStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("3"))
	errorStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1"))

	state := m.engine.State()
	scoreHeader := scoreStyle.Render(fmt.Sprintf("Score: %d | Level: %d", state.Score, state.Level))

	discoveryMessage := ""
	if m.discovered {
		discoveryMessage = "\n" + errorStyle.Render("YOU WERE DISCOVERED! Game reset.") + "\n"
	}

	switch m.mode {
	case "won":
		winStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
		return scoreHeader + "\n\n" + winStyle.Render("CONGRATULATIONS! YOU WON!") + "\n\nYou have successfully completed all missions and puzzles!\n\nPress q or Ctrl+C to quit."
	case "puzzle":
		currentPuzzle := m.engine.NextIncompletePuzzle()
		if currentPuzzle != nil {
			return scoreHeader + discoveryMessage + "\n" + style.Render(currentPuzzle.Prompt) + "\n\nYour answer: " + m.input + "\n\nPress q or Ctrl+C to quit."
		}
		if m.engine.AllMissionsCompleted() && m.engine.AllPuzzlesCompleted() {
			return scoreHeader + "\n\nGame completed! Please restart.\n\nPress q or Ctrl+C to quit."
		} else if !m.engine.AllMissionsCompleted() {
			currentMission := m.engine.NextIncompleteMission()
			if currentMission != nil {
				commands := m.engine.Commands()
				commandList := "Available commands:\n"
				for _, cmd := range commands {
					commandList += "- " + cmd.Name + ": " + cmd.Description + "\n"
				}
				return scoreHeader + discoveryMessage + "\n" + style.Render(currentMission.Title) + "\n" + currentMission.Description + "\n\n" + commandList + "\nType command: " + m.input + "\n\nPress q or Ctrl+C to quit."
			}
		}
		return scoreHeader + "\n\nNo more content available.\n\nPress q or Ctrl+C to quit."
	case "mission":
		currentMission := m.engine.NextIncompleteMission()
		if currentMission != nil {
			commands := m.engine.Commands()
			commandList := "Available commands:\n"
			for _, cmd := range commands {
				commandList += "- " + cmd.Name + ": " + cmd.Description + "\n"
			}

			return scoreHeader + discoveryMessage + "\n" + style.Render(currentMission.Title) + "\n" + currentMission.Description + "\n\n" + commandList + "\nType command: " + m.input + "\n\nPress q or Ctrl+C to quit."
		}
		if m.engine.AllMissionsCompleted() && m.engine.AllPuzzlesCompleted() {
			return scoreHeader + "\n\nGame completed! Please restart.\n\nPress q or Ctrl+C to quit."
		} else if !m.engine.AllPuzzlesCompleted() {
			currentPuzzle := m.engine.NextIncompletePuzzle()
			if currentPuzzle != nil {
				return scoreHeader + discoveryMessage + "\n" + style.Render(currentPuzzle.Prompt) + "\n\nYour answer: " + m.input + "\n\nPress q or Ctrl+C to quit."
			}
		}
		return scoreHeader + "\n\nNo more content available.\n\nPress q or Ctrl+C to quit."
	}

	return scoreHeader + discoveryMessage + "\n" + style.Render("Start your hacking journey!") + "\n\nPress q or Ctrl+C to quit."
}

func main() {
	state := game.NewState(0, 1)
	engine := game.NewEngine(state)
	m := model{
		engine:     engine,
		mode:       "mission",
		input:      "",
		discovered: false,
		gameWon:    false,
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v\n", err)
		os.Exit(1)
	}
}

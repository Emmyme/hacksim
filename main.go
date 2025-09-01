package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Emmyme/hacksim/game"

)

type model struct {
	engine *game.Engine
	lines  []string
}

func (m model) Init() tea.Cmd {return nil}


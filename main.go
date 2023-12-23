package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialGrid())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialGrid() Grid {

	return Grid{
		cells:     map[Position]Cell{},
		computed:  map[Position]string{},
		cursor:    Cursor{Position{row: 1, col: 1}, false, 0, ""},
		selection: []Position{},
	}
}

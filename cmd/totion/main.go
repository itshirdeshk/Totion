package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/itshirdeshk/totion/internal/tui"
)

// Main init function of our go program
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error getting home directory", err)
	}

	tui.VaultDir = fmt.Sprintf("%s/.totion", homeDir)
}

func main() {
	p := tea.NewProgram(tui.InitializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var styleForWelcome = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	var styleForHelp = lipgloss.NewStyle()

	welcome := styleForWelcome.Render("Welcome to Totion 🧠")

	help := styleForHelp.Render("Ctrl+N: new file • Ctrl+L: list • Esc: back • Ctrl+S: save • Ctrl+Q: quit")

	view := ""

	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.noteTextArea.View()
	}

	if m.showingList {
		view = m.list.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help)
}

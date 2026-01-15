package tui

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	noteTextArea           textarea.Model
	list                   list.Model
	showingList            bool
}

var (
	VaultDir string

	textInputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))

	docStyle = lipgloss.NewStyle().
			Margin(1, 2)
)

// Init function of our bubble tea
func (m Model) Init() tea.Cmd {
	return nil
}

func InitializeModel() Model {

	err := os.MkdirAll(VaultDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize new file text input
	ti := textinput.New()
	ti.Placeholder = "Enter file name..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 20
	ti.Cursor.Style = textInputStyle
	ti.PromptStyle = textInputStyle
	ti.TextStyle = textInputStyle

	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.ShowLineNumbers = false
	ta.Focus()

	// note list
	noteList := listFiles()
	finalList := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All List 📙"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("254")).
		Padding(0, 1)

	return Model{
		newFileInput:           ti,
		createFileInputVisible: false,
		noteTextArea:           ta,
		list:                   finalList,
	}
}

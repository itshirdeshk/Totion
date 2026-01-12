package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir       string
	textInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	docStyle       = lipgloss.NewStyle().Margin(1, 2)
)

// Main init function of our go program
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error getting home directory", err)
	}

	vaultDir = fmt.Sprintf("%s/.totion", homeDir)
}

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	noteTextArea           textarea.Model
	list                   list.Model
	showingList            bool
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Init function of our bubble tea
func (m model) Init() tea.Cmd {
	return nil
}

func initializeModel() model {

	err := os.MkdirAll(vaultDir, 0750)
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

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		noteTextArea:           ta,
		list:                   finalList,
	}
}

func (m model) View() string {
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-5)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}
			if m.currentFile != nil {
				m.noteTextArea.SetValue("")
				m.currentFile = nil
			}

			if m.showingList {

				if m.list.FilterState() == list.Filtering {
					break
				}

				m.showingList = false
			}
		case "ctrl+l":
			noteList := listFiles()
			m.list.SetItems(noteList)

			m.showingList = true
			return m, nil
		case "enter":
			if m.currentFile != nil {
				break
			}

			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filePath := fmt.Sprintf("%s/%s", vaultDir, item.title)
					content, err := os.ReadFile(filePath)

					if err != nil {
						log.Printf("error while reading the note: %v", err)
					}

					m.noteTextArea.SetValue(string(content))

					f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("error while reading the note: %v", err)
					}

					m.currentFile = f
					m.showingList = false
				}
				return m, nil

			}

			// todo: create file
			fileName := m.newFileInput.Value()
			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", vaultDir, fileName)

				if _, err := os.Stat(filePath); err == nil {
					return m, nil
				}

				f, err := os.Create(filePath)
				if err != nil {
					log.Fatal(err)
				}
				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")
			}
			return m, nil
		case "ctrl+s":
			if m.currentFile == nil {
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("can not save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("can not save the file :(")
				return m, nil
			}

			noteContent := m.noteTextArea.Value()
			if _, err := m.currentFile.WriteString(noteContent); err != nil {
				fmt.Println("can not save the file :(")
				return m, nil
			}

			err := m.currentFile.Close()
			if err != nil {
				fmt.Println("can not close the file")
			}

			m.currentFile = nil
			m.noteTextArea.SetValue("")
			return m, nil
		}

		if m.createFileInputVisible {
			m.newFileInput, cmd = m.newFileInput.Update(msg)
		}
		if m.currentFile != nil {
			m.noteTextArea, cmd = m.noteTextArea.Update(msg)
		}
		if m.showingList {
			m.list, cmd = m.list.Update(msg)
		}
	}

	return m, cmd
}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func listFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("error reading notes")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			modTime := info.ModTime().Format("2006-01-02 15:04:05")
			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modTime),
			})
		}
	}

	return items
}

package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
					filePath := fmt.Sprintf("%s/%s", VaultDir, item.title)
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
				filePath := fmt.Sprintf("%s/%s.md", VaultDir, fileName)

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

func listFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(VaultDir)
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

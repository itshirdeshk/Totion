# Totion 🧠

A terminal-based note-taking application built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea). Totion provides a beautiful, interactive CLI interface for creating, editing, and managing markdown notes directly from your terminal.

## Features

- **Interactive TUI**: Beautiful terminal user interface powered by Bubble Tea
- **Quick Note Creation**: Create new markdown notes instantly with `Ctrl+N`
- **Note Management**: View and browse all your notes in an organized list
- **File Editing**: Edit notes with a built-in text area editor
- **Auto-Save**: Save your work with `Ctrl+S`
- **Local Storage**: All notes are stored locally in `~/.totion` directory
- **Markdown Support**: All notes are saved as `.md` files

## Installation

### Prerequisites

- Go 1.25.5 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/itshirdeshk/totion.git
cd totion

# Build the application
make build

# Run the application
make run
```

Alternatively, you can build and run manually:

```bash
go build -o totion .
./totion
```

## Usage

### Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Ctrl+N` | Create a new note |
| `Ctrl+L` | List all notes |
| `Ctrl+S` | Save current note |
| `Ctrl+Q` or `Q` | Quit application |
| `Esc` | Go back/Cancel current action |
| `Enter` | Confirm action/Open selected note |

### Workflow

1. **Create a New Note**: Press `Ctrl+N`, enter a filename, and press `Enter`
2. **Write Your Note**: Start typing in the text area
3. **Save Your Note**: Press `Ctrl+S` to save changes
4. **View All Notes**: Press `Ctrl+L` to see a list of all your notes
5. **Open Existing Note**: Select a note from the list and press `Enter`

## Storage

All notes are stored as markdown files in:
- **Linux/macOS**: `~/.totion/`
- **Windows**: `C:\Users\<YourUsername>\.totion\`

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal output

## Project Structure

```
totion/
├── main.go           # Main application logic and TUI implementation
├── go.mod            # Go module dependencies
├── makefile          # Build automation
└── README.md         # Project documentation
```

## Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests

## License

This project is open source. Please check the repository for license details.

## Author

[Hirdesh Khandelwal](https://github.com/itshirdeshk)

---

Built with ❤️ using [Charm](https://github.com/charmbracelet) libraries

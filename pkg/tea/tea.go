package tea

import (
    "os"
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    question string
    answers  []string
    index   int 
}

func initModel() model {
    return model {
        question: "do you like go",
        answers: []string{"yes*", "no i like js (cringe)"},
        index: 0,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func Exec() {
    p := tea.NewProgram(initModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.index > 0 {
                m.index--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.index < len(m.answers)-1 {
                m.index++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
					fmt.Println("pressed enter")
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
	s := m.question + "\n"
	for i, ans := range m.answers{
		cursor := " " // no cursor
		if m.index == i {
				cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, ans)
	}
	return s
}

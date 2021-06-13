package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type inputErr error

type InputHandler func(answer string) error

type modelInput struct {
	input textinput.Model
	err   error

	InputFunc InputHandler
	Question  string
}

func NewInput(placeholder string, handler InputHandler) *tea.Program {
	ti := textinput.NewModel()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.TextStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: "#202020",
			Dark:  "#00d8eb",
		},
	)

	mi := modelInput{
		input:     ti,
		err:       nil,
		InputFunc: handler,
	}

	return tea.NewProgram(mi)
}

func (m modelInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	commands := []tea.Cmd{cmd}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "ctrl+c", "esc":
			err := m.InputFunc(m.input.Value())
			if err != nil {
				return m, tea.Quit
			}
			return m, tea.Quit
		default:
			m.input, _ = m.input.Update(msg)
			return m, nil
		}

	case inputErr:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(commands...)
}

func (m modelInput) View() string {
	output := fmt.Sprintf("%s\n\n%s\n\n%s\n",
		m.Question,
		m.input.View(),
		"(ctrl+c to quit)",
	)

	return output
}

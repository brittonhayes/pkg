package components

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type spinnerErr error

type modelSpinner struct {
	spinner  spinner.Model
	err      error
	quitting bool

	Message string
}

func NewSpinner(message string, color string) *tea.Program {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	model := modelSpinner{
		spinner:  s,
		Message:  message,
	}
	return tea.NewProgram(model)
}

func (m modelSpinner) Init() tea.Cmd {
	return spinner.Tick
}

func (m modelSpinner) quit() (tea.Model, tea.Msg) {
	return m, tea.Quit()
}

func (m modelSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch message := msg.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, nil
		default:
			return m, nil
		}
	case spinnerErr:
		m.err = message
		return m, nil
	}

	if m.quitting {
		m.err = errors.New(m.Message)
		return m, tea.Quit
	}

	var teaCmd tea.Cmd
	m.spinner, teaCmd = m.spinner.Update(msg)
	return m, teaCmd
}

func (m modelSpinner) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s %s\n\n", m.spinner.View(), m.Message)
	return str
}

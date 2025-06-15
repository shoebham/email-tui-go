package model

import (
	"email-client/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type NewMailModel struct {
	textInputs []textinput.Model
	focusIndex int
}

func (m *NewMailModel) Init() tea.Cmd {
	// Initialize the new mail model, if needed
	return textinput.Blink
}

func InitialNewMailModel() *NewMailModel {
	m := NewMailModel{
		textInputs: make([]textinput.Model, 3),
		focusIndex: 0,
	}

	var t textinput.Model

	for i := range m.textInputs {
		t = textinput.New()
		t.CharLimit = 64
		t.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("012"))

		switch i {
		case 0:
			t.Placeholder = "Enter recipient email"
			t.Focus()
			t.Prompt = "To: "
			t.Width = 50

		case 1:
			t.Placeholder = "Enter subject"
			t.Prompt = "Subject: "
			t.Width = 50
			t.CharLimit = 128
		case 2:
			t.Placeholder = "Enter body"
			t.Prompt = "Body: "
			t.Width = 50

		}
		m.textInputs[i] = t
	}
	return &m
}
func (m *NewMailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Cycle focus index
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			if m.focusIndex > len(m.textInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.textInputs)
			}

			cmds := make([]tea.Cmd, len(m.textInputs))
			for i := range m.textInputs {
				if i == m.focusIndex {
					cmds[i] = m.textInputs[i].Focus()
					m.textInputs[i].PromptStyle = utils.FocusedStyle
				} else {
					m.textInputs[i].Blur() // returns updated model and cmd
					m.textInputs[i].PromptStyle = utils.NoStyle
					m.textInputs[i].TextStyle = utils.NoStyle
				}
			}
		case "ctrl+backspace":
			// Handle backspace to go back to the inbox
			return m, func() tea.Msg {
				return SelectedEmailMsg{Email: EmailItem{}}
			}

		}
	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}
func (m *NewMailModel) View() string {
	var s string
	s += utils.TitleStyle.Render("New Mail") + "\n\n"
	// Render each text input
	// and add a newline between them.
	var b strings.Builder
	for i, input := range m.textInputs {
		m.textInputs[i].PlaceholderStyle = utils.PlaceholderStyle
		b.WriteString(input.View())
		if i < len(m.textInputs)-1 {
			b.WriteRune('\n')
		}

	}
	s += b.String() + "\n\n"
	return utils.AppStyle.Render(s)

}

func (m *NewMailModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.textInputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.textInputs {

		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
	}

	for i, input := range m.textInputs {
		if i == m.focusIndex {
			input.PromptStyle = utils.FocusedStyle
			input.CompletionStyle = utils.FocusedStyle
			input.TextStyle = utils.PlaceholderStyle
		} else {
			input.PromptStyle = utils.NoStyle
		}
		m.textInputs[i] = input
	}

	return tea.Batch(cmds...)
}

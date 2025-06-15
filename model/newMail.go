package model

import (
	"email-client/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

const (
	to = iota
	subject
	body
	send
)

type NewMailModel struct {
	textInputs []textinput.Model
	focusIndex int
	body       textarea.Model
}

func (m *NewMailModel) Init() tea.Cmd {
	// Initialize the new mail model, if needed
	return textinput.Blink
}

func InitialNewMailModel() *NewMailModel {
	m := NewMailModel{
		textInputs: make([]textinput.Model, 2),
		focusIndex: 0,
		body:       textarea.New(),
	}

	var t textinput.Model
	var body textarea.Model
	body = textarea.New()
	body.Placeholder = "Enter your message here"

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

		}
		m.textInputs[i] = t
		m.body = body
	}
	return &m
}
func (m *NewMailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab", "up", "down":

			s := msg.String()
			// Cycle focus index
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			if m.focusIndex > send {
				m.focusIndex = to
			}
			if m.focusIndex < to {
				m.focusIndex = send

			}
			if m.focusIndex == body {
				m.body.Focus()
			}

			cmds := make([]tea.Cmd, len(m.textInputs))
			for i := range m.textInputs {
				if i == m.focusIndex {
					cmds[i] = m.textInputs[i].Focus()
					m.textInputs[i].PromptStyle = utils.FocusedStyle
					m.body.Blur()
				} else {
					m.textInputs[i].Blur() // returns updated model and cmd
					m.textInputs[i].PromptStyle = utils.NoStyle
					m.textInputs[i].TextStyle = utils.NoStyle
				}
			}
		case "enter":
			if m.focusIndex == send {
				// Handle sending the email
				to := m.textInputs[0].Value()
				subject := m.textInputs[1].Value()
				body := m.body.Value()

				// Here you would typically send the email using an SMTP client or similar.
				// For now, we will just print it to the console.
				fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)

				// Reset the model after sending
				return InitialNewMailModel(), nil
			} else {
				m.textInputs[m.focusIndex].Blur()
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

	sendButton := utils.BlurredButton
	if m.focusIndex == send {
		sendButton = utils.FocusedButton
		m.body.Blur()
	}
	body := fmt.Sprintf(
		"Body.\n\n%s\n\n%s",
		m.body.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
	s += b.String() + "\n\n"
	s += body
	s += sendButton + "\n\n"
	return utils.AppStyle.Render(s)

}

func (m *NewMailModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.textInputs))
	var cmd tea.Cmd

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.textInputs {
		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
	}
	m.body, cmd = m.body.Update(msg)
	cmds = append(cmds, cmd)
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

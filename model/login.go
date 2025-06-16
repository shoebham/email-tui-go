package model

import (
	"email-client/utils"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	output     string
	Width      int
	Height     int
}

var ()

func InitialLoginModel() *LoginModel {
	m := LoginModel{inputs: make([]textinput.Model, 2)}

	var t textinput.Model

	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = utils.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Enter your email"
			t.Focus()
			t.Prompt = "Email: "
			t.CharLimit = 64
		case 1:
			t.Placeholder = "Enter your password"
			t.Prompt = "Password: "
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = 'x'

		}
		m.inputs[i] = t
	}
	m.focusIndex = 0
	return (&m)
}

func (m *LoginModel) Init() tea.Cmd {

	return textinput.Blink
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				email := m.inputs[0].Value()
				return m, func() tea.Msg {
					return LoginSuccessMsg{Username: email}
				}
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus() // returns a Cmd to set focus
					m.inputs[i].PromptStyle = utils.FocusedStyle
					m.inputs[i].TextStyle = utils.FocusedStyle
				} else {
					m.inputs[i].Blur() // returns updated model and cmd
					m.inputs[i].PromptStyle = utils.NoStyle
					m.inputs[i].TextStyle = utils.NoStyle
				}
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m LoginModel) View() string {
	var b strings.Builder
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}
	button := &utils.BlurredButton
	if m.focusIndex == len(m.inputs) {
		button = &utils.FocusedButton
	}

	var s string
	butt := fmt.Sprintf("\n\n%s\n\n%s", *button, "(ctrl+c to quit)")
	s += b.String() + "\n\n"
	s += butt

	return utils.AppStyle.Width(m.Width).Height(m.Height).Render(s)
	//return b.String()
}

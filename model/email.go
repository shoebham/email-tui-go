package model

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EmailModel struct {
	CurrentItem EmailItem
}
type EmailItem struct {
	subject  string
	body     string
	sender   string
	receiver string
}

func InitialEmailModel(item EmailItem) *EmailModel {
	return &EmailModel{
		CurrentItem: item,
	}
}

func (i EmailItem) Title() string       { return i.subject }
func (i EmailItem) Body() string        { return i.body }
func (i EmailItem) Description() string { return i.body }
func (i EmailItem) FilterValue() string { return i.subject }
func (i EmailItem) Receiver() string    { return i.receiver }
func (i EmailItem) Sender() string      { return i.sender }

func (m *EmailModel) Init() tea.Cmd {
	// Initialize the email model, if needed
	return nil
}
func (m *EmailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "backspace":
			// Handle backspace to go back to the inbox
			return m, func() tea.Msg {
				return SelectedEmailMsg{Email: EmailItem{}}
			}
		}
	}

	return m, nil
}
func (m *EmailModel) View() string {
	if m.CurrentItem == (EmailItem{}) {
		return "No email selected"
	}
	s := selectedEmailView(m.CurrentItem)
	return appStyle.Render(s)
}

func selectedEmailView(item EmailItem) string {
	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#04B575")).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center).
		Padding(1, 2).
		Render
	return border(
		titleStyle.Render(item.Title()) + "\n" +
			lipgloss.NewStyle().Padding(1, 2).Render(
				fmt.Sprintf(
					"Description: %s\nSender: %s\nReceiver: %s",
					item.Body(),
					senderStyle.Render(item.Sender()),
					receiverStyle(item.Receiver()),
				),
			) + "\n" +
			statusMessageStyle("Press 'backspace' to go back to the inbox") + "\n")

}

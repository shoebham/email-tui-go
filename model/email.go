package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type emailModel struct {
	currentItem emailItem
}
type emailItem struct {
	subject  string
	body     string
	sender   string
	receiver string
}

func newEmailModel(item emailItem) emailModel {
	return emailModel{
		currentItem: item,
	}
}

func (i emailItem) Title() string       { return i.subject }
func (i emailItem) Body() string        { return i.body }
func (i emailItem) Description() string { return i.body }
func (i emailItem) FilterValue() string { return i.subject }
func (i emailItem) Receiver() string    { return i.receiver }
func (i emailItem) Sender() string      { return i.sender }

func (m emailModel) Init() tea.Cmd {
	// Initialize the email model, if needed
	return nil
}
func (m emailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m emailModel) View() string {
	if m.currentItem == nil {
		return "No email selected"
	}
	s := selectedEmailView(m.currentItem.(emailItem))
	return appStyle.Render(s)
}

func selectedEmailView(item emailItem) string {
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
					senderStyle(item.Sender()),
					receiverStyle(item.Receiver()),
				),
			) + "\n" +
			statusMessageStyle("Press 'backspace' to go back to the inbox") + "\n")

}

package model

import (
	"email-client/utils"
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
	return utils.AppStyle.Render(s)
}

func selectedEmailView(item EmailItem) string {
	emailStyle := utils.SelectedEmailStyle
	senderPaddingLeft := (100 - len("Sender: "+item.Sender()) - 2)
	receiverPaddingLeft := (100 - len("Receiver: "+item.Receiver())) - 2
	return emailStyle.Padding(1).Render(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s\n%s",
			utils.SelectedEmailSubjectStyle.Render(item.Title()),
			utils.SenderStyle.PaddingLe(senderPaddingLeft).Render("Sender: "+item.Sender()),
			utils.ReceiverStyle.PaddingLeft(receiverPaddingLeft).Render("Receiver: "+item.Receiver()),
			utils.SelectedEmailBodyStyle.Render(item.Body()),
		),
	) + "\n" + lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render("Press 'backspace' to return to inbox")

}

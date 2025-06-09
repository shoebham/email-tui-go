package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InboxModel struct {
	mails          []EmailItem
	emailListModel []emailListModel
	currentIdx     int
	selected       EmailItem
}

type inboxMsg struct {
	mails []EmailItem
}

var (
	appStyle   = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#25A065")).
			Padding(1, 2)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	senderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
			Padding(0, 1).
			Render

	receiverStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9A9A9")).
			Padding(0, 1).
			Render

	// Styling for the email list items
	normalItemStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#25A065")).
			Width(100).
			Height(3)

	selectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#FFFFFF")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FFFDF5")).
				Width(100).
				Height(3)

	subjectStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFDF5"))

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9A9A9"))
)

func fetchEmails() tea.Msg {
	// Simulate fetching emails
	emails := []string{
		"Email 1: Welcome to our service!",
		"Email 2: Your account has been created.",
		"Email 3: Don't forget to verify your email.",
	}
	items := make([]EmailItem, len(emails))
	for i, email := range emails {
		items[i] = EmailItem{
			subject:  email,
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		}
	}

	return inboxMsg{mails: items}
}

func InitialInboxModel() *InboxModel {

	// Create initial items
	items := []EmailItem{
		EmailItem{
			subject:  "Email 1: Welcome to our service!",
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		},
		EmailItem{
			subject:  "Email 2: Your account has been created.",
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		},
		EmailItem{
			subject:  "Email 3: Don't forget to verify your email.",
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		},
	}

	mails := make([]EmailItem, len(items))
	return &InboxModel{
		mails:    mails,
		selected: EmailItem{},
		emailListModel: []emailListModel{{
			sender:    "shubham",
			subject:   "Welcome to our service!",
			shortBody: "This is a sample email body.",
		}, {
			sender:    "shubham",
			subject:   "Welcome to our service!",
			shortBody: "This is a sample email body.",
		}, {
			sender:    "shubham",
			subject:   "Welcome to our service!",
			shortBody: "This is a sample email body.",
		}, {
			sender:    "shubham",
			subject:   "Welcome to our service!",
			shortBody: "This is a sample email body.",
		}, {
			sender:    "shubham",
			subject:   "Welcome to our service!",
			shortBody: "This is a sample email body.",
		},
		},
	}

}

func (m *InboxModel) Init() tea.Cmd {

	return fetchEmails
}
func (m *InboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.currentIdx <= len(m.mails) {
				m.currentIdx++
			}
		case "k", "up":
			if m.currentIdx > 0 {
				m.currentIdx--
			}
		case "enter":
			if m.currentIdx <= len(m.mails) {
				selectedEmail := m.mails[m.currentIdx]
				m.selected = selectedEmail
			}
			return m, func() tea.Msg {
				return SelectedEmailMsg{Email: m.selected}
			}
		case "backspace":
			if m.selected != (EmailItem{}) {
				m.selected = EmailItem{}
			}
		}
	}

	if inboxMsg, ok := msg.(inboxMsg); ok {
		m.mails = inboxMsg.mails
		if len(m.mails) > 0 {
			m.currentIdx = 0 // Reset current index to the first email
		} else {
			m.currentIdx = -1 // No emails available
		}
	}
	return m, tea.Batch(cmds...)
}

func (m *InboxModel) View() string {
	var s string

	if m.currentIdx == -1 {
		m = InitialInboxModel()
	}
	s += titleStyle.Render("📧 Inbox") + "\n\n"

	for i, email := range m.emailListModel {
		style := normalItemStyle
		if i == m.currentIdx {
			style = selectedItemStyle
		}

		content := fmt.Sprintf("%s\n%s\n%s",
			subjectStyle.Render(email.subject),
			senderStyle(email.sender),
			bodyStyle.Render(email.shortBody),
		)

		s += style.Render(content) + "\n"
	}

	return appStyle.Render(s)
}

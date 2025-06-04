package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InboxModel struct {
	mails    list.Model
	selected EmailItem
}

type inboxMsg struct {
	mails []list.Item
}

var (
	appStyle   = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

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
)

func fetchEmails() tea.Msg {
	// Simulate fetching emails
	emails := []string{
		"Email 1: Welcome to our service!",
		"Email 2: Your account has been created.",
		"Email 3: Don't forget to verify your email.",
	}
	items := make([]list.Item, len(emails))
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
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#04B575"))
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#04B575"))

	// Create initial items
	items := []list.Item{
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

	mails := list.New(items, delegate, 0, 0)
	mails.Title = titleStyle.Render("Inbox")
	mails.SetShowStatusBar(false)
	mails.SetFilteringEnabled(false)

	return (&InboxModel{
		mails:    mails,
		selected: EmailItem{},
	})
}

func (m *InboxModel) Init() tea.Cmd {
	// Set list styles
	m.mails.Styles.Title = titleStyle

	return fetchEmails
}
func (m *InboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.mails.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			selectedEmail := m.mails.SelectedItem()
			if email, ok := selectedEmail.(EmailItem); ok {
				return m, func() tea.Msg {
					return SelectedEmailMsg{Email: email}
				}
			}
		case "backspace":
			if m.selected != (EmailItem{}) {
				m.selected = EmailItem{}
			}
		}
	}

	var cmd tea.Cmd
	m.mails, cmd = m.mails.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *InboxModel) View() string {
	var s string
	s = m.mails.View()
	return appStyle.Render(s) + "\n"
}

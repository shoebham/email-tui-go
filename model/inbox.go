package model

import (
	"email-client/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type InboxModel struct {
	mails      []EmailItem
	currentIdx int
	selected   EmailItem
}

type inboxMsg struct {
	mails []EmailItem
}

// create a enum of colors that i can access

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
			if m.currentIdx < len(m.mails)-1 {
				m.currentIdx++
			}
		case "k", "up":
			if m.currentIdx > 0 {
				m.currentIdx--
			}
		case "enter":
			if m.currentIdx < len(m.mails) {
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
	s += utils.TitleStyle.MarginLeft(50).Render("Inbox") + "\n\n"
	for i, email := range m.mails {
		style := utils.NormalItemStyle

		sender := utils.SenderStyle.Render(email.sender)
		subject := utils.SubjectStyle.PaddingLeft(2).Render(email.subject)
		if i == m.currentIdx {
			style = utils.SelectedItemStyle
			sender = utils.SelectedSenderStyle.Render(email.sender)
			subject = utils.SelectedSubjectStyle.PaddingLeft(2).Render(email.subject)
		}
		content := fmt.Sprintf("%s%s", sender, subject)
		s += style.Render(content) + "\n"
	}

	return utils.AppStyle.Render(s)
}

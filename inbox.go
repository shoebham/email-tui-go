package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type inboxModel struct {
	mails    list.Model
	selected list.Item
}

type emailModel struct {
	subject  string
	body     string
	sender   string
	receiver string
	date     string
}
type emailMsg struct {
	email emailModel
}

type inboxMsg struct {
	mails []list.Item
}

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

func fetchEmails() tea.Msg {
	fmt.Printf("Clled")
	// Simulate fetching emails
	emails := []string{
		"Email 1: Welcome to our service!",
		"Email 2: Your account has been created.",
		"Email 3: Don't forget to verify your email.",
	}

	items := make([]list.Item, len(emails))
	for i, email := range emails {
		items[i] = item{
			title:       email,
			description: "This is a sample email description.",
		}
	}

	return inboxMsg{mails: items}
}

func initialModel() inboxModel {
	m := inboxModel{}

	// Create a list model
	mails := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	mails.Title = titleStyle.Render("Inbox")
	mails.SetShowStatusBar(false)
	mails.SetFilteringEnabled(false)

	// Set the initial items
	mails.SetItems([]list.Item{})

	m.mails = mails

	return m
}

func (m inboxModel) Init() tea.Cmd {
	return func() tea.Msg {
		return fetchEmails()
	}
}

func (m inboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.selected = m.mails.SelectedItem()
			selectedEmail := m.mails.SelectedItem().(item)
			emailDetails := emailModel{
				subject:  selectedEmail.title,
				body:     selectedEmail.description,
				sender:   "x",
				receiver: "y",
			}
			return updateEmailView(msg, m, emailDetails)
		case "backspace":
			if m.selected != nil {
				m.selected = nil
			} else {
				return m, nil
			}

		}

	}
	// This will also call our delegate's update function.
	newListModel, _ := m.mails.Update(msg)
	m.mails = newListModel
	if inboxMsg, ok := msg.(inboxMsg); ok {
		m.mails.SetItems(inboxMsg.mails)

	}

	return m, tea.Batch(cmds...)
}

func (m inboxModel) View() string {
	var s string
	if m.selected != nil {
		s = selectedEmailView(m.mails.SelectedItem().(item))
	} else {
		s = m.mails.View()
	}

	return appStyle.Render(s) + "\n"
}

func selectedEmailView(item item) string {
	return fmt.Sprintf(
		"Selected Email:\nTitle: %s\nDescription: %s\n",
		item.Title(),
		item.Description(),
	)

}

func updateEmailView(msg tea.Msg, m inboxModel, email emailModel) (tea.Model, tea.Cmd) {
	fmt.Sprintf("Subject: %s\nBody: %s\nSender: %s\nReceiver: %s\n",
		email.subject, email.body, email.sender, email.receiver)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
		if msg.String() == "h" {
			m.selected = nil
			return m, nil
		}
	}
	return m, nil
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

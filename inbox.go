package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type inboxModel struct {
	mails    list.Model
	selected emailItem
}

type emailModel struct {
	subject  string
	body     string
	sender   string
	receiver string
	date     string
}

type inboxMsg struct {
	mails []list.Item
}

type emailItem struct {
	subject  string
	body     string
	sender   string
	receiver string
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
	senderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
			Padding(0, 1).
			Render

	// light greyed color for receiver
	receiverStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9A9A9")).
			Padding(0, 1).
			Render
)

func (i emailItem) Title() string       { return i.subject }
func (i emailItem) Body() string        { return i.body }
func (i emailItem) Description() string { return i.body }
func (i emailItem) FilterValue() string { return i.subject }
func (i emailItem) Receiver() string    { return i.receiver }
func (i emailItem) Sender() string      { return i.sender }

func fetchEmails() tea.Msg {
	// Simulate fetching emails
	emails := []string{
		"Email 1: Welcome to our service!",
		"Email 2: Your account has been created.",
		"Email 3: Don't forget to verify your email.",
	}

	items := make([]list.Item, len(emails))
	for i, email := range emails {
		items[i] = emailItem{
			subject:  email,
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		}
	}

	return inboxMsg{mails: items}
}
func initialModel() inboxModel {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#04B575"))
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#04B575"))

	// Create initial items
	items := []list.Item{
		emailItem{
			subject:  "Email 1: Welcome to our service!",
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		},
		emailItem{
			subject:  "Email 2: Your account has been created.",
			body:     "This is a sample email body.",
			sender:   "shubham",
			receiver: "shubham",
		},
		emailItem{
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

	return inboxModel{
		mails:    mails,
		selected: emailItem{},
	}
}

func (m inboxModel) Init() tea.Cmd {
	// Set list styles
	m.mails.Styles.Title = titleStyle

	// Return the command directly
	return fetchEmails
}
func (m inboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.mails.SetSize(msg.Width-h, msg.Height-v)

	case inboxMsg:
		// Handle incoming emails first
		items := msg.mails
		m.mails.SetItems(items)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if i := m.mails.SelectedItem(); i != nil {
				m.selected = i.(emailItem)
			}
		case "backspace":
			if m.selected != (emailItem{}) {
				m.selected = emailItem{}
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

func (m inboxModel) View() string {
	var s string
	if m.selected != (emailItem{}) {
		s = selectedEmailView(m.mails.SelectedItem().(emailItem))
	} else {
		s = m.mails.View()
	}

	return appStyle.Render(s) + "\n"
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

//
//func updateEmailView(msg tea.Msg, m inboxModel) (tea.Model, tea.Cmd) {
//
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "q", "ctrl+c":
//			return m, tea.Quit
//		}
//		if msg.String() == "h" {
//			m.selected = nil
//			return m, nil
//		}
//	}
//	return m, nil
//}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type inboxModel struct {
	mails list.Model
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
	return appStyle.Render(m.mails.View())
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

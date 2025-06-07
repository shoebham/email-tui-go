package model

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type emailListModel struct {
	sender    string
	subject   string
	shortBody string
}

var (
	boldTextStyle        = lipgloss.NewStyle().Bold(true)
	fadedTextStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#A9A9A9"))
	highlightedTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	normalTextStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	dividerStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#A9A9A9")).Render
	normalBorderStyle    = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
)

func (e *emailListModel) Init() tea.Cmd {

	// Initialize the model, if needed
	e = &emailListModel{
		sender:    "shubham",
		subject:   "Welcome to our service!",
		shortBody: "This is a sample email body.",
	}
	return func() tea.Msg {
		return e
	}
}

func (e *emailListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return e, tea.Quit
		}
	}
	return e, nil
}

func (e *emailListModel) View() string {

	e.sender = boldTextStyle.Render(e.sender)
	e.subject = highlightedTextStyle.Render(e.subject)
	e.shortBody = fadedTextStyle.Render(e.shortBody)
	// Format the email item for display
	s := fmt.Sprintf("%s\t%s\t%s", e.sender, e.subject, e.shortBody)

	return lipgloss.NewStyle().
		Width(100).
		Height(1).
		Border(lipgloss.RoundedBorder()).
		Render(s)

}

package utils

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Background(lipgloss.AdaptiveColor{Light: White, Dark: Black})

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	SenderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: Black, Dark: Gray})

	SubjectStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
			Background(lipgloss.AdaptiveColor{Light: White, Dark: Black})

	ReceiverStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9A9A9"))

	bottomBorder = lipgloss.Border{
		Bottom: "─",
	}
	// Styling for the email list items
	NormalItemStyle = lipgloss.NewStyle().
			Background(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
			Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
			Border(bottomBorder).
			BorderForeground().
			Width(100)

	SelectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
				Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
				Width(100)

	SelectedEmailBorder = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Background(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				BorderForeground(lipgloss.AdaptiveColor{Light: Black, Dark: White})
	SelectedEmailSubjectStyle = lipgloss.NewStyle().
					Bold(true).
					Border(bottomBorder).
					BorderBackground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
					BorderForeground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
					Background(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
					Foreground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
					AlignHorizontal(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Width(97)

	SelectedEmailBodyStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
				Background(lipgloss.AdaptiveColor{Light: White, Dark: Black})
	SelectedEmailStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Background(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.AdaptiveColor{Light: Black, Dark: White}).
				Width(100)
	SelectedSenderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Background(lipgloss.AdaptiveColor{Light: Black, Dark: White})

	SelectedSubjectStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
				Background(lipgloss.AdaptiveColor{Light: Black, Dark: White})
)

var FocusedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
	Foreground(lipgloss.Color("012"))

var BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

var CursorStyle = FocusedStyle

var NoStyle = lipgloss.NewStyle()

var HelpStyle = BlurredStyle

var CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

var FocusedButton = FocusedStyle.Render("[ Submit ]")
var BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
var PlaceholderStyle = lipgloss.NewStyle().
	Background(lipgloss.AdaptiveColor{Light: White, Dark: Black}).
	Foreground(lipgloss.AdaptiveColor{Light: Gray, Dark: Gray})

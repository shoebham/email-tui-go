package utils

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"os"
)

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(uintptr(int(os.Stdin.Fd())))
	return width, height, err
}

func GetWindowMsgTypeForInbox() tea.Msg {
	width, height, err := GetTerminalSize()
	if err != nil {
		// fallback to default size if error
		width, height = 80, 24
	}
	msg := tea.WindowSizeMsg{
		Width:  width,
		Height: height,
	}
	return msg
}

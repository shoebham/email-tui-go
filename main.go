package main

import (
	"email-client/model"

	tea "github.com/charmbracelet/bubbletea"
)

type page int

const (
	loginPage page = iota
	inboxPage
	emailPage
)

type rootModel struct {
	currentPage page
	inbox       *model.InboxModel
	login       *model.LoginModel
	email       *model.EmailModel
}

func (m *rootModel) Init() tea.Cmd {
	switch m.currentPage {
	case loginPage:
		return m.login.Init()
	case inboxPage:
		return m.inbox.Init()
	case emailPage:
		return m.email.Init()
	default:
		return nil
	}
}

func (m *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {

	case model.SelectedEmailMsg:
		if msg.Email == (model.EmailItem{}) {
			m.currentPage = inboxPage

		} else {
			m.email = &model.EmailModel{
				CurrentItem: msg.Email,
			}
			m.currentPage = emailPage
		}
		return m, nil

	case model.LoginSuccessMsg:
		m.currentPage = inboxPage
		return m, m.inbox.Init()

	}

	switch m.currentPage {
	case emailPage:
		var cmd tea.Cmd
		updatedModel, cmd := m.email.Update(msg)
		m.email = updatedModel.(*model.EmailModel)
		if cmd != nil {
			return m, cmd

		}
		return m, nil

	case inboxPage:
		var cmd tea.Cmd
		updatedModel, cmd := m.inbox.Update(msg)
		m.inbox = updatedModel.(*model.InboxModel)
		return m, cmd
	case loginPage:
		var cmd tea.Cmd
		updatedModel, cmd := m.login.Update(msg)
		m.login = updatedModel.(*model.LoginModel)
		return m, cmd
	default:
		panic("unhandled default case")
	}
	return m, nil
}

func (m *rootModel) View() string {
	switch m.currentPage {
	case inboxPage:
		return m.inbox.View()
	case loginPage:
		return m.login.View()
	case emailPage:
		return m.email.View()
	default:
		panic("unhandled default case")
	}
}

func main() {
	p := tea.NewProgram(&rootModel{
		currentPage: loginPage,
		inbox:       model.InitialInboxModel(),
		login:       model.InitialLoginModel(),
		email:       model.InitialEmailModel(model.EmailItem{}),
	}, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if err := p.Start(); err != nil {
		panic(err)
	}

}

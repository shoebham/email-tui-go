package main

import (
	"email-client/model"
	"email-client/utils"
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
		m.email = &model.EmailModel{
			CurrentItem: msg.Email,
		}
		m.currentPage = emailPage
		return m, nil
	case model.LoginSuccessMsg:
		m.currentPage = inboxPage
		m.Update(utils.GetWindowMsgTypeForInbox())
	}
	switch m.currentPage {
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
	})

	if err := p.Start(); err != nil {
		panic(err)
	}

}

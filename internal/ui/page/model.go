package page

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Model, tea.Cmd)
	View() string
	GetPageTitle() string
}

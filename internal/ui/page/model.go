package page

import tea "charm.land/bubbletea/v2"

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Model, tea.Cmd)
	View() string
	GetPageTitle() string
}

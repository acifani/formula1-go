package program

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Help   key.Binding
	Quit   key.Binding
	Select key.Binding
	WDC    key.Binding
	WCC    key.Binding
	Season key.Binding
	Back   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Season, k.WDC, k.WCC, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Season, k.WDC, k.WCC},
		{k.Up, k.Down, k.Select},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	WCC: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "constructor standings"),
	),
	WDC: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "driver standings"),
	),
	Season: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "current season"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
}

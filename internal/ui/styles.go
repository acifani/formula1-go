package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	borderColor = lipgloss.Color("240")
	purple      = lipgloss.Color("57")
	cream       = lipgloss.Color("229")
)

type Styles struct {
	Wrap, Paragraph, Base, Title, Footer lipgloss.Style
	Table, SelectableTable               table.Styles
}

func NewStyles() Styles {
	s := Styles{}
	s.Wrap = lipgloss.NewStyle().Width(58)
	s.Paragraph = s.Wrap.Margin(1, 0, 1, 2)
	s.Title = lipgloss.NewStyle().Margin(1, 0, 0, 1).Padding(0, 1).Background(purple).Bold(true)
	s.Footer = lipgloss.NewStyle().Margin(1, 0, 0, 1)

	s.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Margin(1, 0, 0, 0)

	s.Table = table.DefaultStyles()
	s.Table.Header = s.Table.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		BorderBottom(true)
	s.Table.Selected = s.Table.Selected.
		UnsetForeground().
		UnsetBackground().
		Bold(false)

	s.SelectableTable = s.Table
	s.SelectableTable.Selected = s.SelectableTable.Selected.
		Foreground(cream).
		Background(purple)

	return s
}

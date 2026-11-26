package ui

import "charm.land/bubbles/v2/table"

// TableHeaderHeight is the number of lines the table header takes: the
// title row plus its bottom border. table.SetHeight counts these lines.
const TableHeaderHeight = 2

// NewTable builds a table as wide as its columns. A table with no width
// renders no rows.
func NewTable(columns []table.Column, styles table.Styles) table.Model {
	width := 0
	for _, c := range columns {
		width += c.Width + styles.Cell.GetHorizontalFrameSize()
	}

	return table.New(
		table.WithColumns(columns),
		table.WithStyles(styles),
		table.WithWidth(width),
	)
}

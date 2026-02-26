package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MyTable(pages *tview.Pages) *tview.Table {
	// ===== Таблица =====
	table := tview.NewTable().SetBorders(true)
	table.SetBorder(true).SetTitle(" Таблица ")
	headers := []string{"ID", "Имя", "Статус"}
	for i, h := range headers {
		table.SetCell(0, i, tview.NewTableCell(h).SetSelectable(false))
	}
	data := [][]string{
		{"1", "Server01", "Online"},
		{"2", "Server02", "Offline"},
		{"3", "Server03", "Online"},
	}
	for r, row := range data {
		for c, col := range row {
			table.SetCell(r+1, c, tview.NewTableCell(col))
		}
	}
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("menu")
			return nil
		}
		return event
	})
	return table
}

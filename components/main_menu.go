package components

import "github.com/rivo/tview"

func MainMenu(app *tview.Application, pages *tview.Pages, title string) tview.Primitive {
	menu := tview.NewList().
		AddItem("Форма ввода", "input + password", '1', func() {
			pages.SwitchToPage("form")
		}).
		AddItem("Таблица", "Таблица", '2', func() {
			pages.SwitchToPage("table")
		}).
		AddItem("MsgBox", "Сообщение", '5', func() {
			pages.SwitchToPage("msg")
		}).
		AddItem("Выход", "Закрыть программу", '0', func() {
			pages.SwitchToPage("exitmodal")
		})

	menu.SetBorder(true).SetTitle(" " + title + " ").SetTitleAlign(tview.AlignLeft)
	return menu
}

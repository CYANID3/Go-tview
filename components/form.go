package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MainForm(pages *tview.Pages) *tview.Form {
	form := tview.NewForm().
		AddDropDown("Пол", []string{"  Мужской  ", "  Женский  "}, 0, nil).
		AddInputField("Имя", "", 10, nil, nil).
		AddInputField("Мымя", "", 10, nil, nil).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorWhite).SetLabelColor(tcell.ColorBlue)
	// name := form.GetFormItemByLabel("Имя").(*tview.InputField).GetText()

	form.AddButton("OK", func() {
		pages.SwitchToPage("menu")
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage("menu")
	})

	form.SetButtonsAlign(tview.AlignLeft)
	form.SetButtonBackgroundColor(tcell.ColorRed)

	form.SetBorder(true).SetTitle(" Форма ").SetTitleAlign(tview.AlignLeft)

	return form
}

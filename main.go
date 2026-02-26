package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	// ===== Лог-окно =====
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	logView.SetBorder(true).SetTitle(" Логи ")

	addLog := func(text string) {
		fmt.Fprintln(logView, text)
	}

	// ===== Главное меню =====
	menu := tview.NewList().
		AddItem("Форма ввода", "input + password", '1', func() {
			pages.SwitchToPage("form")
		}).
		AddItem("Подтверждение", "Yes/No окно", '2', func() {
			pages.SwitchToPage("confirm")
		}).
		AddItem("Таблица", "Пример таблицы", '3', func() {
			pages.SwitchToPage("table")
		}).
		AddItem("InputBox", "Окно ввода текста", '4', func() {
			pages.SwitchToPage("inputbox")
		}).
		AddItem("MsgBox", "Сообщение", '5', func() {
			pages.SwitchToPage("msgbox")
		}).
		AddItem("Логи", "Открыть лог-окно", '6', func() {
			pages.SwitchToPage("logs")
		}).
		AddItem("Выход", "Закрыть программу", 'q', func() {
			app.Stop()
		})
	menu.SetBorder(true).SetTitle(" Главное меню ")

	// ===== Форма =====
	form := tview.NewForm().
		AddDropDown("Пол", []string{"Мужской", "Женский"}, 0, nil).
		AddInputField("Имя", "", 10, nil, nil).
		AddPasswordField("Пароль", "", 10, '*', nil)

	form.AddButton("OK", func() {
		name := form.GetFormItemByLabel("Имя").(*tview.InputField).GetText()
		addLog("Форма отправлена, имя: " + name)
		pages.SwitchToPage("menu")
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage("menu")
	})
	form.SetBorder(true).SetTitle(" Форма ").SetTitleAlign(tview.AlignLeft)

	// ===== Подтверждение =====
	confirm := tview.NewModal().
		SetText("Вы уверены?").
		AddButtons([]string{"Да", "Нет"}).
		SetDoneFunc(func(i int, label string) {
			addLog("Нажата кнопка: " + label)
			if label == "Да" {
				app.Stop()
			} else {
				pages.SwitchToPage("menu")
			}
		})

	// ===== MsgBox =====
	msgbox := tview.NewModal().
		SetText("Операция выполнена успешно").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(i int, label string) {
			pages.SwitchToPage("menu")
		})

	// ===== InputBox =====
	inputForm := tview.NewForm().
		AddInputField("Введите текст", "", 30, nil, nil)
	inputForm.AddButton("OK", func() {
		text := inputForm.GetFormItemByLabel("Введите текст").(*tview.InputField).GetText()
		addLog("Введено: " + text)
		pages.SwitchToPage("menu")
	})
	inputForm.AddButton("Отмена", func() {
		pages.SwitchToPage("menu")
	})
	inputForm.SetBorder(true).SetTitle(" InputBox ")

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

	// ===== Лог-экран =====
	logLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logView, 0, 1, false).
		AddItem(tview.NewButton("Назад").SetSelectedFunc(func() {
			pages.SwitchToPage("menu")
		}), 1, 0, true)
	logLayout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("menu")
			return nil
		}
		return event
	})

	// Подсказка для формы
	hint := tview.NewTextView().
		SetText("Shift - Следующий | Shift + Tab - Предыдущий | Esc - меню").
		SetTextColor(tcell.ColorYellow).
		SetTextAlign(tview.AlignCenter)

	// Flex для формы
	formFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true). // Форма занимает всё пространство
		AddItem(hint, 1, 1, false) // Подсказка снизу

	// ===== Pages =====
	pages.AddPage("menu", menu, true, true)
	pages.AddPage("form", form, true, false)
	pages.AddPage("confirm", confirm, true, false)
	pages.AddPage("msgbox", msgbox, true, false)
	pages.AddPage("inputbox", inputForm, true, false)
	pages.AddPage("table", table, true, false)
	pages.AddPage("logs", logLayout, true, false)
	pages.AddPage("form", formFlex, true, false)

	// ===== Глобальный обработчик Esc =====
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("menu")
			return nil
		}
		return event
	})

	// ===== Запуск =====
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"time"

	cmp "project/components"
	cfg "project/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	// ===== Debounce + глобальное отслеживание ESC =====
	cfg.App.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev == nil {
			return ev
		}

		// ===== Debounce =====
		if (ev.Key() == cfg.LastKey && ev.Rune() == cfg.LastRune) && time.Since(cfg.LastTime) < cfg.Debounce {
			return nil
		}
		cfg.LastKey = ev.Key()
		cfg.LastRune = ev.Rune()
		cfg.LastTime = time.Now()

		// ===== Глобальный Esc =====
		if ev.Key() == tcell.KeyEsc {
			cfg.Pages.SwitchToPage("menu")
			return nil
		}

		return ev
	})

	// ===== Создание страниц =====
	menu := cmp.MainMenu(cfg.App, cfg.Pages, "")
	table := cmp.MyTable(cfg.Pages)
	form := cmp.MainForm(cfg.Pages)
	hint := cmp.MyHint(cfg.Pages)
	yesnomodal := cmp.YesNoModal(cfg.Pages, "Точно выйти?", "ДА", "НЕТ", "Success", cfg.App.Stop)
	msg := cmp.CustomMessage(cfg.Pages, "Успех", "ОК", "Warning", func() { cfg.Pages.SwitchToPage("menu") })

	// Flex для формы
	formFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true). // Форма занимает всё пространство
		AddItem(hint, 1, 1, false) // Подсказка снизу

	cfg.Pages.AddPage("menu", menu, true, true)
	cfg.Pages.AddPage("form", formFlex, true, false)
	cfg.Pages.AddPage("yesnomodal", yesnomodal, true, false)
	cfg.Pages.AddPage("msg", msg, true, false)
	cfg.Pages.AddPage("table", table, true, false)

	// ===== Запуск =====
	if err := cfg.App.SetRoot(cfg.Pages, true).Run(); err != nil {
		panic(err)
	}
}

package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MyHint(pages *tview.Pages) *tview.TextView {
	hint := tview.NewTextView().
		SetText("Shift - Следующий | Shift + Tab - Предыдущий | Esc - меню").
		SetTextColor(tcell.ColorYellow).
		SetTextAlign(tview.AlignCenter)
	return hint
}

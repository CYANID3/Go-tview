package components

import (
	"github.com/rivo/tview"
)

func YesNoModal(pages *tview.Pages, question string, positive string, negative string, myFunc func()) *tview.Modal {
	modal := tview.NewModal().
		SetText(question).
		AddButtons([]string{positive, negative}).
		SetDoneFunc(func(i int, label string) {
			if label == positive {
				myFunc()
			} else {
				pages.SwitchToPage("menu")
			}
		})
	return modal
}

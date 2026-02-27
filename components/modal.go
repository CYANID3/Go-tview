package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func modalStyleTyper(t string) {
	switch t {
	case "Danger":
		textColor = "white"
		borderColor = tcell.ColorWhite
		bgColor = tcell.ColorRed
	case "Success":
		textColor = "white"
		borderColor = tcell.ColorWhite
		bgColor = tcell.ColorGreen
	case "Warning":
		textColor = "black"
		borderColor = tcell.ColorBlack
		bgColor = tcell.ColorOrange
	default:
		textColor = "white"
		borderColor = tcell.ColorWhite
		bgColor = tcell.ColorBlue
	}
}

func YesNoModal(pages *tview.Pages, question string, positive string, negative string, modaltype string, myFunc func()) *tview.Modal {
	modalStyleTyper(modaltype)
	modal := tview.NewModal().
		SetText("[" + textColor + "]" + question + "[-]").SetBackgroundColor(bgColor).
		AddButtons([]string{"[grey]" + positive + "[-]", "[grey]" + negative + "[-]"}).
		SetDoneFunc(func(i int, label string) {
			if label == positive {
				myFunc()
			} else {
				pages.SwitchToPage("menu")
			}
		})
	modal.SetBorder(true).SetBorderColor(borderColor).SetBackgroundColor(bgColor)
	return modal
}

package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	textColor   = "white"
	borderColor = tcell.ColorBlack
	bgColor     = tcell.ColorBlue
)

func bgTyper(t string) {
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

func CustomMessage(pages *tview.Pages, message string, btntext string, msgtype string, myFunc func()) *tview.Modal {
	bgTyper(msgtype)
	msgbox := tview.NewModal().
		SetText("[" + textColor + "]" + message + "[-]").SetBackgroundColor(bgColor).
		AddButtons([]string{"[black]" + btntext + "[-]"}).
		SetDoneFunc(func(i int, label string) {
			myFunc()
		})
	msgbox.SetBorder(true).SetBorderColor(borderColor).SetBackgroundColor(bgColor)
	return msgbox
}

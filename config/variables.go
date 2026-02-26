package config

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	// ===== Основные переменные для UI =====
	App   = tview.NewApplication()
	Pages = tview.NewPages()
	// ===== Переменные для debounce =====
	LastKey  tcell.Key
	LastRune rune
	LastTime time.Time
	Debounce = 200 * time.Millisecond
)

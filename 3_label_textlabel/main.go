package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	// New app
	a := app.New()
	// New Window
	w := a.NewWindow("Here is my title")

	// Resize Window
	w.Resize(fyne.NewSize(300, 300))
	// Our First widget
	labelA := widget.NewLabel("Here is first Label")
	w.SetContent(labelA)
	w.ShowAndRun()
}
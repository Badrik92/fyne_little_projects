package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	// creating new app 
	a := app.New()
	// New window
	w := a.NewWindow("Here is my titile")
	w.Resize(fyne.NewSize(400, 400))

	// BUTTON
	btn := widget.NewButton("Press me", func() {
		// Our is ready
		fmt.Println("Here is Go Button")
	})
	// Using our button
	w.SetContent(btn)
	// Runing app
	w.ShowAndRun()
}
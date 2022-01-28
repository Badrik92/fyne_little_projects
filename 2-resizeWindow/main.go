package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
func main() {
	a := app.New()
	w := a.NewWindow("My title window")  
	// Resizing fyne app window (width, heigth)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	// New Window and titile
	w := a.NewWindow("Here title")
	w.Resize(fyne.NewSize(400, 400))

	// Check box widget
	checkbox1 := widget.NewCheck("True", func(b bool) {
		if b == true {
			fmt.Println("Print true")
		}else{
			fmt.Println("Print false")
		}
	})

	// set up content
	w.SetContent(checkbox1)
	w.ShowAndRun()
}
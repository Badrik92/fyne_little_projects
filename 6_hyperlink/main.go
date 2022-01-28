package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Here is widget")

	w.Resize(fyne.NewSize(400, 400))
	// Creating url
	url,_ := url.Parse("https:/blogvali.com")

	// hyperLinc widget
	// first value is Label
	// 2nd value is URL/ websiye adress
	link := widget.NewHyperlink("Visit me", url)
	w.SetContent(link)
	w.ShowAndRun()
}
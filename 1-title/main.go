package main

// Import fyne
import(
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	
	w := a.NewWindow("My first title")

	w.ShowAndRun()
}
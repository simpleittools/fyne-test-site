package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type BlogPost struct {
	title    string
	date     string
	content  string
	expanded bool
}

func main() {
	a := app.New()
	// This should not be needed anymore
	//a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Ryan's Blog")
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

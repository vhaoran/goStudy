package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

/*
# in debian
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
*/

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetTitle("my first demo")
	w.SetContent(widget.NewHBox(
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),

			widget.NewCheck("checkbox", nil),
			widget.NewProgressBarInfinite(),
			widget.NewProgressBar(),
			widget.NewRadio([]string{"a", "b", "c"}, nil),
		),
		widget.NewVBox(
			widget.NewMultiLineEntry(),
		),

		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

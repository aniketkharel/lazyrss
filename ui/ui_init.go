package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func init_ui() {
	a :=app.New()
	w := a.NewWindow("lazyreader")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButtin("Hi", func() {
			hello.setText("Welcome bruh")
		}),
	))
	w.ShowAndRun()
}

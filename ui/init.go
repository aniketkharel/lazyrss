package ui

import (
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func Init_Ui() {
	// app := app.New()
	// window := app.NewWindow("lazyrss")
	//
	// list := widget.NewList(
	// 	func() int {
	// 		return len(data)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("Just a template")
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(data[i])
	// 	})
	//
	// window.SetContent(list)
	// window.ShowAndRun()
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}

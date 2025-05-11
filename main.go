package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/aniketkharel/rssreader/ui"
	"github.com/aniketkharel/rssreader/workers"
)

func main() {
	workers.Start_scrapping()
	go func() {
		window := new(app.Window)
		window.Option(app.Size(340, 600))
		window.Option(app.Title("lazyrss"))
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			ui.Construct_List(theme, gtx)
			e.Frame(gtx.Ops)
		}
	}
}

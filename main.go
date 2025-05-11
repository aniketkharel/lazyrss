package main

import (
	"log"
	"os"
	"sync"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/aniketkharel/rssreader/models"
	"github.com/aniketkharel/rssreader/ui"
	"github.com/aniketkharel/rssreader/workers"
)

func main() {
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
	data := make(chan models.FeedXML)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		workers.Start_scrapping(data)
	}()
	go func() {
		wg.Wait()
		close(data)
	}()
	var feedList []models.FeedXML
	for feed := range data {
		feedList = append(feedList, feed)
	}
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			ui.Construct_List(theme, gtx, feedList)
			e.Frame(gtx.Ops)
		}
	}
}

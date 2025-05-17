package main

import (
	"sync"
	"time"

	"github.com/aniketkharel/rssreader/models"
	"github.com/aniketkharel/rssreader/ui"
	"github.com/aniketkharel/rssreader/workers"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

func main() {
	//fmt.Println(len(os.Args), os.Args)
	run()
}

func run() {
	s.Start()
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
	s.Stop()
	for v := range data {
		// items = v.Channel.Items
		color.Green(ui.Format_Channel_Info(v))
	}
}

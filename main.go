package main

import (
	"github.com/aniketkharel/rssreader/models"
	"github.com/aniketkharel/rssreader/ui"
	"github.com/aniketkharel/rssreader/workers"
	"github.com/briandowns/spinner"
	"sync"
	"time"
)

var s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

func main() {
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
		s.Start()
		items := v.Channel.Items
		ui.Format_FeedItem_Info(&items)
		s.Stop()
	}
}

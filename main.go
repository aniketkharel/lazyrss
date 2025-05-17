package main

import (
	"github.com/aniketkharel/rssreader/models"
	"github.com/aniketkharel/rssreader/workers"
	"sync"
)

func main() {
	run()
}

func run() {
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
	for _, v := range feedList {
		println(v.Channel.Title)
	}
}

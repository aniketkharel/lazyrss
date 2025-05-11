package main

import (
	"github.com/aniketkharel/rssreader/ui"
	"github.com/aniketkharel/rssreader/workers"
)

func main() {
	workers.Start_scrapping()
	ui.Init_Ui()
}

package workers

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aniketkharel/rssreader/filesys"
	"github.com/aniketkharel/rssreader/models"
)

func Start_scrapping(data chan models.FeedXML) {
	var wg sync.WaitGroup
	configs := filesys.Load_config()
	if len(configs.Urls) > 0 {
		for _, v := range configs.Urls {
			wg.Add(1)
			go go_to_url(v, &wg, data)
		}
		wg.Wait()
	}
	return
}

func go_to_url(url string, wg *sync.WaitGroup, data chan models.FeedXML) {
	defer wg.Done()
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var feeds models.FeedXML
	err = xml.Unmarshal(body, &feeds)
	if err != nil {
		log.Fatal(err)
	}
	data <- feeds
}

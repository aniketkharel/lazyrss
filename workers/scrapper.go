package workers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aniketkharel/rssreader/models"
)

func Start_scrapping() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.9news.com.au/victoria/rss",
		"https://news.ycombinator.com/rss",
	}

	for _, v := range urls {
		wg.Add(1)
		go go_to_url(v, &wg)
	}
	wg.Wait()
}

func go_to_url(url string, wg *sync.WaitGroup) {
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var feeds models.FeedXML
	err = xml.Unmarshal(body, &feeds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feeds)
	defer wg.Done()
}

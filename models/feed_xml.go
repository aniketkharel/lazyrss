package models

type FeedItemXML struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

type FeedXML struct {
	Channel struct {
		Title       string        `xml:"title"`
		Link        string        `xml:"link"`
		Description string        `xml:"description"`
		Items       []FeedItemXML `xml:"item"`
	} `xml:"channel"`
}

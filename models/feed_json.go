package models

type FeedItemJSON struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	PubDate     string `json:"pubDate"`
	Description string `json:"description"`
}

type FeedJSON struct {
	Channel struct {
		Title       string         `json:"title"`
		Link        string         `json:"link"`
		Description string         `json:"description"`
		Items       []FeedItemJSON `json:"item"`
	} `json:"channel"`
}

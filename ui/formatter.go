package ui

import (
	"fmt"

	"github.com/aniketkharel/rssreader/models"
)

func Format_Channel_Info(feed models.FeedXML) string {
	return fmt.Sprintf("%s - %s\n", feed.Channel.Title, feed.Channel.Link)
}

func Format_FeedItem_Info(item models.FeedItemXML) string {
	return fmt.Sprintf("%s \n %s %s %s\n", item.Title, item.Description,
		item.Description, item.Link,
	)
}

package ui

import (
	"fmt"
	"github.com/aniketkharel/rssreader/models"
	"github.com/fatih/color"
	"golang.org/x/net/html"
	"strings"
)

func Format_FeedItem_Info(items *[]models.FeedItemXML) {
	for i, v := range *items {
		fmt.Println(color.HiYellowString("------------------------------------------------------------"))
		fmt.Printf("%s\n", color.HiGreenString("Feed #%d", i+1))
		fmt.Printf("%s: %s\n", color.HiCyanString("Title"), v.Title)
		fmt.Printf("%s: %s\n", color.HiBlueString("Link"), v.Link)
		fmt.Printf("%s: %s\n", color.HiWhiteString("Published"), v.PubDate)
		desc := renderDescription(v.Description)
		fmt.Printf("%s:\n%s\n", color.HiMagentaString("Description"), desc)
	}
	fmt.Println(color.HiYellowString("============================================================"))
}

func renderDescription(htmlStr string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return htmlStr
	}
	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return b.String()
}

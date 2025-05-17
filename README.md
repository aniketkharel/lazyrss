# Lazyrss

Probably, lazy as me. Trying to get up and running with a minimal terminal RSS
reader. As this is my first ever project with go language, I am missing on lot
of features, hopefully in near future, I will improve the application. (maybe!).

## Setup

```sh
git clone https://github.com/aniketkharel/lazyrss.git
cd lazyrss
go run main.go
```

## Configuration

By deafault, when you lunch the program, it will create a **config.json** file inside **HOME_DIR/.config/lazyrss** folder. Then you can update the list of urls inside json file to fetch the **RSS Feeds**.

```json
{
  "urls": [
    "https://www.9news.com.au/victoria/rss",
    "https://news.ycombinator.com/rss",
    "https://www.xda-developers.com/feed/tag/rss/",
    "https://techcrunch.com/feed/"
  ]
}
```

As of now, I have tested the program in my work machine, running Fedora so, there is no guarantee that this will work in Window OS. By the way, this is meant for _tmux_ and _terminal_ power users. HAHA ! :D

## References

1. [go-colorable](github.com/mattn/go-colorable)
2. [colors](https://github.com/fatih/color.git)
3. [spinners](github.com/briandowns/spinner)

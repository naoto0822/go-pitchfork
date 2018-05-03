package pitchfork

import (
	"testing"

	"github.com/naoto0822/gss/gss"
)

type testRssImpl struct {
	Items map[string][]gss.Item
}

func newTestRss() rss {
	items := make(map[string][]gss.Item)

	thumbnail := gss.Thumbnail{
		URL:    "http://hoge/foo.png",
		Width:  90,
		Height: 90,
	}

	author := gss.Author{
		Name:  "Win Butler",
		Email: "arcadefire@gmail.com",
	}
	authors := []Author{author}

	categories := []string{"news", "music"}

	// empty: Content, Image, Enclosure,,,
	item := gss.Item{
		ID:          "abcdefghijk",
		Title:       "Arcade Fire new release",
		Link:        "https://arcadefire.com/",
		Description: "Everything Now",
		PubDate:     "Sat, 02 Apr 2017 05:18:00 +0000",
		Updated:     "Tue, 10 May 2017 13:40:16 +0000",
		Authors:     authors,
		Categories:  categories,
		Thumbnail:   thumbnail,
	}

	// for news
	items["https://pitchfork.com/rss/news"] = []gss.Item{item}

	// for not found
	items["notfound"] = []gss.Item{}

	return &testRssImpl{Items: items}
}

func (t *testRssImpl) fetch(url string) ([]gss.Item, error) {
	if items, ok := t.Items[url]; ok {
		return items, nil
	}

	return nil, fmt.Errorf("failed rss fetch")
}

func TestNewRss(t *testing.T) {
	rss := newRss()
	if rss == nil {
		t.Error("TestNewRss: failed factory rss")
	}
}

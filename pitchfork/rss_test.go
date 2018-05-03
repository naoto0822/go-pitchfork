package pitchfork

import (
	"reflect"
	"testing"
	"time"

	"github.com/naoto0822/gss/gss"
)

// testRssClient mock for gss.Client
type testRssClient struct {
	Items map[string][]gss.Item
}

// newTestRssClient create test data
func newTestRssClient() rssClient {
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

	return testRssClient{Items: items}
}

func (t *testRssClient) Parse(url string) (*gss.Result, error) {
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

func TestFetchRss(t *testing.T) {
	testRssClient := newTestRssClient()
	converter := newConverter()
	testRss := rss{
		rssClient: testRssClient,
		converter: converter,
	}

	articles, err := testRss.Fetch("https://pitchfork.com/rss/news")
	if err != nil {
		t.Error("TestFetchRss: not expect return err: ", err)
	}

	if len(articles) != 1 {
		t.Error("TestFetchRss: empty articles")
	}

	rawPubDate := "Sat, 02 Apr 2017 05:18:00 +0000"
	pubDate, _ := time.Parse(time.RFC1123Z, rawPubDate)

	rawUpdated := "Tue, 10 May 2017 13:40:16 +0000"
	updated, _ := time.Parse(time.RFC1123Z, rawUpdated)

	expect := Article{
		ID:          "abcdefghijk",
		Title:       "Arcade Fire new release",
		Link:        "https://arcadefire.com/",
		Description: "Everything Now",
		RawPubDate:  rawPubDate,
		PubDate:     pubDate,
		RawUpdated:  rawUpdated,
		Updated:     updated,
		Author:      "Win Butler",
		Categories:  []string{"news", "music"},
		Thumbnail: Thumbnail{
			URL:    "http://hoge/foo.png",
			Width:  90,
			Height: 90,
		},
	}

	actual := articles[0]
	if !reflect.DeepEqual(expect, actual) {
		t.Error("TestFetchRss: not match expected article")
	}
}

func TestNotFoundFetchRss(t *testing.T) {
	testRssClient := newTestRssClient()
	converter := newConverter()
	testRss := rss{
		rssClient: testRssClient,
		converter: converter,
	}

	articles, err := testRss.Fetch("notfound")
	if err != nil {
		t.Error("TestNotFoundFetchRss: not expect return err: ", err)
	}

	if len(articles) != 0 {
		t.Error("TestNotFoundFetchRss: expect return empty articles")
	}
}

func TestErrorFetchRss(t *testing.T) {
	testRssClient := newTestRssClient()
	converter := newConverter()
	testRss := rss{
		rssClient: testRssClient,
		converter: converter,
	}

	_, err := testRss.Fetch("hoge")
	if err == nil {
		t.Error("TestErrorFetchRss: expect not return articles")
	}
}

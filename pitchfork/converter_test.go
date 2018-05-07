package pitchfork

import (
	"reflect"
	"testing"
	"time"

	"github.com/naoto0822/gss/gss"
)

func TestNewConverter(t *testing.T) {
	c := newConverter()
	if c == nil {
		t.Error("TestNewConverter: faild factory converter")
	}
}

func TestConvertArticle(t *testing.T) {
	thumbnail := gss.Thumbnail{
		URL:    "http://hoge/foo.png",
		Width:  90,
		Height: 90,
	}

	author := gss.Author{
		Name:  "james",
		Email: "lcd@gmail.com",
	}
	authors := []gss.Author{author}

	categories := []string{"news", "music"}

	// empty: Content, Image, Enclosure,,,
	item := gss.Item{
		ID:          "abcdefghijk",
		Title:       "LCD is came back",
		Link:        "https://lcdsoundsystem.com/",
		Description: "LCD SoundSystem is reunion in 2016",
		PubDate:     "Sat, 28 Apr 2018 05:18:00 +0000",
		Updated:     "Tue, 01 May 2018 13:40:16 +0000",
		Authors:     authors,
		Categories:  categories,
		Thumbnail:   thumbnail,
	}

	c := newConverter()
	actual := c.convertArticle(item)

	rawPubDate := "Sat, 28 Apr 2018 05:18:00 +0000"
	pubdate, _ := time.Parse(time.RFC1123Z, rawPubDate)

	rawUpdated := "Tue, 01 May 2018 13:40:16 +0000"
	updated, _ := time.Parse(time.RFC1123Z, rawUpdated)

	expect := Article{
		ID:          "abcdefghijk",
		Title:       "LCD is came back",
		Link:        "https://lcdsoundsystem.com/",
		Description: "LCD SoundSystem is reunion in 2016",
		RawPubDate:  rawPubDate,
		PubDate:     pubdate,
		RawUpdated:  rawUpdated,
		Updated:     updated,
		Author:      "james",
		Categories:  []string{"news", "music"},
		Thumbnail: Thumbnail{
			URL:    "http://hoge/foo.png",
			Width:  90,
			Height: 90,
		},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Error("TestConvertArticle: not match converted Article expect:", expect, ", actual:", actual)
	}
}

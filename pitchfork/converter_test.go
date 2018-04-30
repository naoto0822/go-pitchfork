package pitchfork

import (
	"reflect"
	"testing"

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
	authors := []Author{author}

	categories := []string{"news", "music"}

	// empty: Content, Image, Enclosure,,,
	item := gss.Item{
		ID:          "abcdefghijk",
		Title:       "LCD is came back",
		Link:        "https://lcdsoundsystem.com/",
		Description: "LCD SoundSystem is reunion in 2016",
		PubDate:     "",
		Updated:     "",
		Authors:     authors,
		Categories:  categories,
		Thumbnail:   thumbnail,
	}

	c := newConverter()
	actual := c.convertArticle(item)

	expect := Article{
		ID:          "abcdefghijk",
		Title:       "LCD is came back",
		Link:        "https://lcdsoundsystem.com/",
		Description: "LCD SoundSystem is reunion in 2016",
		RawPubDate:  "",
		RawUpdated:  "",
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

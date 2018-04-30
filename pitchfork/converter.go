package pitchfork

import (
	"github.com/naoto0822/gss/gss"
)

type converter struct{}

func newConverter() *converter {
	return &converter{}
}

// convert from `gss.Item` to `pitchfork.Article`
func (c *converter) convertArticle(from gss.Item) Article {
	thumbnail := Thumbnail{
		URL:    from.Thumbnail.URL,
		Width:  from.Thumbnail.Width,
		Height: from.Thumbnail.Height,
	}

	pubDate, _ := parseTime(i.RawPubDate)
	updated, _ := parseTime(i.RawUpdated)

	var author string
	if len(i.Authors) > 0 {
		a := i.Authors[0]
		author = a.Name
	}

	article := Article{
		ID:          from.ID,
		Title:       from.Title,
		Link:        from.Link,
		Description: from.Description,
		RawPubDate:  i.RawPubDate,
		PubDate:     pubDate,
		RawUpdated:  i.RawUpdated,
		Updated:     updated,
		Author:      author,
		Categories:  i.Categories,
	}
	return article
}

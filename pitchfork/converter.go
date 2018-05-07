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

	pubDate, _ := parseTime(from.PubDate)
	updated, _ := parseTime(from.Updated)

	var author string
	if len(from.Authors) > 0 {
		a := from.Authors[0]
		author = a.Name
	}

	article := Article{
		ID:          from.ID,
		Title:       from.Title,
		Link:        from.Link,
		Description: from.Description,
		RawPubDate:  from.PubDate,
		PubDate:     pubDate,
		RawUpdated:  from.Updated,
		Updated:     updated,
		Author:      author,
		Categories:  from.Categories,
		Thumbnail:   thumbnail,
	}
	return article
}

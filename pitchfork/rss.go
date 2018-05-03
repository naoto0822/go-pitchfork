package pitchfork

import (
	"github.com/naoto0822/gss/gss"
)

// interface for gss.Client
type rssClient interface {
	Parse(string) (*gss.Result, error)
}

type rss struct {
	rssClient rssClient
	converter *converter
}

func newRss() *rss {
	rss := gss.NewClient()
	converter := newConverter()

	return &rss{
		rssClient: rss,
		converter: converter,
	}
}

func (r *rss) fetch(url string) ([]Article, error) {
	ret, err := r.rssClient.Parse(url)
	if err != nil {
		return nil, err
	}

	// access gss.Feed.Items ([]Item)
	if len(ret.Feed.Items) < 1 {
		return nil, ErrorNotFound
	}

	var articles []Article
	for _, i := range ret.Feed.Items {
		a := r.converter.convertArticle(i)
		articles = append(articles, a)
	}

	return articles, nil
}

package pitchfork

import (
	"github.com/naoto0822/gss/gss"
)

type rss interface {
	fetch(string) ([]gss.Item, error)
}

type rssImpl struct {
	rss *gss.Client
}

func newRss() rss {
	rss := gss.NewClient()

	return &rssImpl{
		rss: rss,
	}
}

func (r *rssImpl) fetch(url string) ([]gss.Item, error) {
	ret, err := r.rss.Parse(url)
	if err != nil {
		return nil, err
	}
	// gss.Result.Feed.Items
	return ret.Feed.Items, nil
}

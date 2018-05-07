package pitchfork

import ()

type articleRepository interface {
	fetch(string) ([]Article, error)
}

type articleRepositoryImpl struct {
	rss       rss
	converter *converter
}

func newArticleRepository() articleRepository {
	rss := newRss()
	converter := newConverter()

	return &articleRepositoryImpl{
		rss:       rss,
		converter: converter,
	}
}

func (r *articleRepositoryImpl) fetch(url string) ([]Article, error) {
	// return []gss.Item, error
	items, err := r.rss.fetch(url)
	if err != nil {
		return nil, err
	}

	if len(items) < 1 {
		return nil, ErrorNotFound
	}

	var articles []Article
	for _, i := range items {
		a := r.converter.convertArticle(i)
		articles = append(articles, a)
	}

	return articles, nil
}

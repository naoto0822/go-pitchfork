package pitchfork

import (
	"github.com/naoto0822/gss/gss"
)

type testArticleRepositoryImpl struct {
	articles map[string][]Article
}

func newTestArticleRepository() articleRepository {
	articles := make(map[string][]Article)

	// TODO: factory articles data

	return *testArticleRepositoryImpl{
		articles: articles,
	}
}

func (r *testArticleRepositoryImpl) fetch(url string) ([]Article, error) {
	if target, ok := r.articles[url]; ok {
		return target, nil
	}
	return nil, fmt.Errorf("failed fetch articles")
}

func TestFetchArticle(t *testing.T) {
	rss := newTestRss()
	converter := newConverter()
	repo := articleRepositoryImpl{
		rss:       rss,
		converter: converter,
	}

	articles, err := repo.fetch("https://pitchfork.com/rss/news")
	if err != nil {
		t.Error("TestFetchArticle: not expect return err: ", err)
	}

	if len(articles) != 1 {
		t.Error("TestFetchArticle: empty articles")
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
		t.Error("TestFetchArticle: not match expected article")
	}
}

func TestNotFoundFetchArticle(t *testing.T) {
	rss := newTestRss()
	converter := newConverter()
	repo := articleRepositoryImpl{
		rss:       rss,
		converter: converter,
	}

	articles, err := repo.fetch("notfound")
	if err != nil {
		t.Error("TestNotFoundFetchArticle: not expect return err: ", err)
	}

	if len(articles) != 0 {
		t.Error("TestNotFoundFetchArticle: expect return empty articles")
	}
}

func TestErrorFetchArticle(t *testing.T) {
	rss := newTestRss()
	converter := newConverter()
	repo := articleRepositoryImpl{
		rss:       rss,
		converter: converter,
	}

	_, err := repo.fetch("hoge")
	if err == nil {
		t.Error("TestErrorFetchArticle: expect not return articles")
	}
}

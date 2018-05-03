package pitchfork

// NewsService cf. https://pitchfork.com/rss/news/
type NewsService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewNewsService factory NewsService
func NewNewsService(articleRepository articleRepository, builder *urlBuilder) *NewsService {
	return &NewsService{
		articleRepository: articleRepository,
		urlBuilder:        builder,
	}
}

// Fetch News Feed
func (s *NewsService) Fetch() ([]Article, error) {
	url, err := s.urlBuilder.build(newsType)
	if err != nil {
		return nil, err
	}

	articles, err := s.articleRepository.fetch(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

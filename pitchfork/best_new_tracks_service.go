package pitchfork

// BestNewTracksService cf. https://pitchfork.com/rss/reviews/best/tracks/
type BestNewTracksService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewBestNewTracksService factory BestNewTracksService
func NewBestNewTracksService(a articleRepository, u *urlBuilder) *BestNewTracksService {
	return &BestNewTracksService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Best New Tracks Feed
func (s *BestNewTracksService) Fetch() ([]Article, error) {
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

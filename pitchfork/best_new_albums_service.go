package pitchfork

// BestNewAlbumsService cf. https://pitchfork.com/rss/reviews/best/albums/
type BestNewAlbumsService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewBestNewAlbumsService factory BestNewAlbumsService
func NewBestNewAlbumsService(a articleRepository, u *urlBuilder) *BestNewAlbumsService {
	return &BestNewAlbumsService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Best New Albums Feed
func (s *BestNewAlbumsService) Fetch() ([]Article, error) {
	url, err := s.urlBuilder.build(bestNewAlbumsType)
	if err != nil {
		return nil, err
	}

	articles, err := s.articleRepository.fetch(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

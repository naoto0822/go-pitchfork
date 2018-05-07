package pitchfork

// BestNewReissuesService cf. https://pitchfork.com/rss/reviews/best/reissues/
type BestNewReissuesService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewBestNewReissuesService factory BestNewReissuesService
func NewBestNewReissuesService(a articleRepository, u *urlBuilder) *BestNewReissuesService {
	return &BestNewReissuesService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Best New Reissues Feed
func (s *BestNewReissuesService) Fetch() ([]Article, error) {
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

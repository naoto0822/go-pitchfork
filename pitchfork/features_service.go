package pitchfork

// FeaturesService
// cf. https://pitchfork.com/rss/features/
type FeaturesService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewFeaturesService factory FeaturesService
func NewFeaturesService(a articleRepository, u *urlBuilder) *FeaturesService {
	return &FeaturesService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Features Feed
func (s *FeaturesService) Fetch() ([]Article, error) {
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

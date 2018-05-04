package pitchfork

// PitchService
// cf. https://pitchfork.com/rss/thepitch/
type PitchService struct {
	articleRepository articleRepository
	urlBuilder        urlBuilder
}

// NewPitchService factory PitchService
func NewPitchService(a articleRepository, u urlBuilder) *PitchService {
	return &PitchService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Pitch
func (r *PitchService) Fetch() ([]Article, error) {
	url, err := s.urlBuilder.build(pitchType)
	if err != nil {
		return nil, err
	}

	articles, err := s.articleRepository.fetch(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

package pitchfork

// TrackReviewsService
// cf. https://pitchfork.com/rss/reviews/tracks/
type TrackReviewsService struct {
	articleRepository articleRepository
	urlBuilder        urlBuilder
}

// NewTrackReviewsService factory TrackReviewsService
func NewTrackReviewsService(a articleRepository, u urlBuilder) TrackReviewsService {
	return &TrackReviewsService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Track Reviews
func (r *TrackReviewsService) Fetch() ([]Article, error) {
	url, err := s.urlBuilder.build(trackReviewsType)
	if err != nil {
		return nil, err
	}

	articles, err := s.articleRepository.fetch(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

package pitchfork

// AlbumReviewsService cf. https://pitchfork.com/rss/reviews/albums/
type AlbumReviewsService struct {
	articleRepository articleRepository
	urlBuilder        *urlBuilder
}

// NewAlbumReviewsService factory AlbumReviewsService
func NewAlbumReviewsService(a articleRepository, u *urlBuilder) *AlbumReviewsService {
	return &AlbumReviewsService{
		articleRepository: a,
		urlBuilder:        u,
	}
}

// Fetch Album Reviews
func (s *AlbumReviewsService) Fetch() ([]Article, error) {
	url, err := s.urlBuilder.build(albumReviewsType)
	if err != nil {
		return nil, err
	}

	articles, err := s.articleRepository.fetch(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

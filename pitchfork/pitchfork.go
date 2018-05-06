package pitchfork

import ()

// Client is communication PitchFork RSS
type Client struct {
	News            *NewsService
	AlbumReviews    *AlbumReviewsService
	TrackReviews    *TrackReviewsService
	Pitch           *PitchService
	Features        *FeaturesService
	BestNewAlbums   *BestNewAlbumsService
	BestNewReissues *BestNewReissuesService
	BestNewTracks   *BestNewTracksService

	// TODO: later implements and attachable HTTP header
	// userAgent string
}

// NewClient factory pitchfork.Client
func NewClient() *Client {
	repo := newArticleRepository()
	urlBuilder := newURLBuilder()

	news := NewNewsService(repo, urlBuilder)
	albumReviews := NewAlbumReviewsService(repo, urlBuilder)
	trackReviews := NewTrackReviewsService(repo, urlBuilder)
	pitch := NewPitchService(repo, urlBuilder)
	features := NewFeaturesService(repo, urlBuilder)
	bestNewAlbums := NewBestNewAlbumsService(repo, urlBuilder)
	bestNewReissues := NewBestNewReissuesService(repo, urlBuilder)
	bestNewTracks := NewBestNewTracksService(repo, urlBuilder)

	return &Client{
		News:            NewsService,
		AlbumReviews:    albumReviews,
		TrackReviews:    trackReviews,
		Pitch:           pitch,
		Features:        features,
		BestNewAlbums:   bestNewAlbums,
		BestNewReissues: bestNewReissues,
		BestNewTracks:   bestNewTracks,
	}
}

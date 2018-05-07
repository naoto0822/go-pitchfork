package pitchfork

import (
	"net/url"
)

const (
	baseURL = "https://pitchfork.com/rss/"
)

type urlBuilder struct {
	BaseURL *url.URL
}

func newURLBuilder() *urlBuilder {
	u, _ := url.Parse(baseURL)
	return &urlBuilder{
		BaseURL: u,
	}
}

func (u *urlBuilder) build(t feedType) (string, error) {
	path := u.getPath(t)
	if path == "" {
		return "", ErrorUnknownType
	}

	url, _ := u.BaseURL.Parse(path)
	return url.String(), nil
}

func (u *urlBuilder) getPath(t feedType) string {
	switch t {
	case newsType:
		return "news/"
	case albumReviewsType:
		return "reviews/albums/"
	case trackReviewsType:
		return "reviews/tracks/"
	case pitchType:
		return "thepitch/"
	case featuresType:
		return "features/"
	case bestNewAlbumsType:
		return "reviews/best/albums/"
	case bestNewReissuesType:
		return "reviews/best/reissues/"
	case bestNewTracksType:
		return "reviews/best/tracks/"
	default:
		return ""
	}
}

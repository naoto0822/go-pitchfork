package pitchfork

import ()

// Client is communication PitchFork RSS
type Client struct {
	baseURL string
	rss     *gss.Client

	// TODO: later
	userAgent string
}

// NewClient factory pitchfork.Client
func NewClient() *Client {
	c := Client{}

	// TODO:

	return &c
}

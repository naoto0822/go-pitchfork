package pitchfork

import (
	"errors"
)

var (
	// ErrorNotFound if rss client returns empty Item slice
	ErrorNotFound = errors.New("go-pitchfork: not found pitchfork.Articles")
	// ErrorUnknownType if using unknown `feedType`
	ErrorUnknownType = errors.New("go-pitchfork: unknown feed type")
)

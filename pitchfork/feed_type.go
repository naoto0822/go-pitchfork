package pitchfork

// feedType is PitchFork Feed types
// cf. https://pitchfork.com/rss/
type feedType int

const (
	// News Feed
	newsType feedType = iota
	// Album Reviews Feed
	albumReviewsType
	// Track Reviews Feed
	trackReviewsType
	// The Pitch Feed
	pitchType
	// Features Feed
	featuresType
	// Best New Album Feed
	bestNewAlbumType
	// Best New Reissues Feed
	bestNewReissuesType
	// Best New Track Feed
	bestNewTrackType
)

// String implmented `fmt.Stringer`
func (f feedType) String() string {
	switch f {
	case newsType:
		return "News Feed"
	case albumReviewsType:
		return "Album Reviews Feed"
	case trackReviewsType:
		return "Track Reviews Feed"
	case pitchType:
		return "The Pitch Feed"
	case featuresType:
		return "Features Feed"
	case bestNewAlbumType:
		return "Best New Album Feed"
	case bestNewReissuesType:
		return "Best New Reissues Feed"
	case bestNewTrackType:
		return "Best New Track Feed"
	default:
		return "Unknown Feed"
	}
}

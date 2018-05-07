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
	bestNewAlbumsType
	// Best New Reissues Feed
	bestNewReissuesType
	// Best New Track Feed
	bestNewTracksType
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
	case bestNewAlbumsType:
		return "Best New Album Feed"
	case bestNewReissuesType:
		return "Best New Reissues Feed"
	case bestNewTracksType:
		return "Best New Track Feed"
	default:
		return "Unknown Feed"
	}
}

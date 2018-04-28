package pitchfork

import (
	"time"
)

// PitchFork RSS date format example: Sat, 28 Apr 2018 05:18:00 +0000
// following doc is written "the Date and Time Specification of RFC 822", but actually it is RFC 1123 format.
// cf. https://cyber.harvard.edu/rss/rss.html
//
// RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700"
// cf. https://golang.org/pkg/time/#pkg-constants
func parseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC1123Z, s)
}

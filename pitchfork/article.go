package pitchfork

import (
	"time"
)

/*

Example: News Feed

<item>
  <title>Vampire Weekend Return at Surprise Show</title>
  <description>For HQ Trivia host Scott Rogowsky’s Running Late event in L.A.</description>
  <link>https://pitchfork.com/news/vampire-weekend-return-at-surprise-show</link>
  <guid isPermaLink="false">5ae0e28ca85afb79a8445d2e</guid>
  <pubDate>Wed, 25 Apr 2018 21:15:34 +0000</pubDate>
  <media:content/>
  <category>News</category>
  <dc:creator>Sam Sodomsky</dc:creator>
  <dc:modified>Thu, 26 Apr 2018 14:53:59 +0000</dc:modified>
  <dc:publisher>Cond Nast</dc:publisher>
  <media:thumbnail url="https://media.pitchfork.com/photos/5ae0e6c08bec5b23c213a369/master/pass/Vampire_Weekend.jpg" width="790" height="395"/>
</item>

*/

// Article <item> tag in feed
type Article struct {
	// ID <guid> tag
	ID string
	// Title <title> tag
	Title string
	// Link <link> tag
	Link string
	// Description <description> tag
	Description string
	// RawPubDate <pubDate> tag
	RawPubDate string
	// PubDate is converted `RawPubDate`
	PubDate time.Time
	// RawUpdated <dc:modified> tag
	RawUpdated string
	// Updated is converted `RawUpdated`
	Updated time.Time
	// Author <dc:creator> tag
	Author string
	// Categories <category> tag
	Categories []string
	// Thumbnail <media:thumbnail> tag in feed
	Thumbnail Thumbnail
}

// Thumbnail <media:thumbnail> tag in feed
type Thumbnail struct {
	// URL attr url
	URL string
	// Width attr width
	Width int64
	// Height attr height
	Height int64
}

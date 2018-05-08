# go-pitchfork

go-pitchfork is Go library for fetching Pitchfork RSS Feed.   
for details see the following.  
- [Pitchfork RSS](https://pitchfork.com/rss/)

## Usage

```go
import "github.com/naoto0822/go-pitchfork/pitchfork"
```

Construct a new pitchfork client, then use the services on the client to access RSS feed.
For example: fetch news feed.

```go
c := pitchfork.NewClient()
articles, err := c.News.Fetch()
```

## Features

- [ ] integration test
- [ ] write doc.go

## License

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.

# go-pitchfork

[![Build Status](https://travis-ci.org/naoto0822/go-pitchfork.svg?branch=master)](https://travis-ci.org/naoto0822/go-pitchfork)
[![GoDoc](https://godoc.org/github.com/naoto0822/go-pitchfork/pitchfork?status.svg)](https://godoc.org/github.com/naoto0822/go-pitchfork/pitchfork)
[![Go Report Card](https://goreportcard.com/badge/github.com/naoto0822/go-pitchfork)](https://goreportcard.com/report/github.com/naoto0822/go-pitchfork)
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/naoto0822/go-pitchfork/blob/master/LICENSE)

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
- [ ] post travis-ci to slack

## License

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.

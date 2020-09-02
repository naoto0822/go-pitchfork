你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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

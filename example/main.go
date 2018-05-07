package main

import (
	"fmt"

	"github.com/naoto0822/go-pitchfork/pitchfork"
)

func main() {
	fmt.Println("go-pitchfork!")

	// News Feed
	c := pitchfork.NewClient()
	articles, err := c.News.Fetch()
	if err != nil {
		fmt.Errorf("failed fetch News Feed err:", err)
	}
	for _, a := range articles {
		fmt.Printf("%+v\n", a)
	}
}

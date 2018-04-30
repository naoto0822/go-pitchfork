package pitchfork

import (
	"testing"
)

func TestNewURLBuilder(t *testing.T) {
	b := newURLBuilder()
	if b == nil {
		t.Error("TestNewURLBuilder: failed factory urlBuilder")
	}
}

func TestBuildURL(t *testing.T) {
	b := newURLBuilder()
	if b == nil {
		t.Error("TestBuildURL: failed factory urlBuilder")
	}

	actual, err := b.build(newsType)
	if err != nil {
		t.Error("TestBuildURL: not expected error: ", err)
	}

	expect := "https://pitchfork.com/rss/news/"
	if actual != expect {
		t.Error("TestBuildURL: not expected url: ", expect)
	}

	// TODO:
	actual, err = b.build(10)
	if err != ErrorUnknownType {
		t.Error("TestBuildURL: not expected error: ", err)
	}

	expect = ""
	if actual != expect {
		t.Error("TestBuildURL: not expected url: ", expect)
	}
}

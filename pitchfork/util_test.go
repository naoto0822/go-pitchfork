package pitchfork

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	// valid
	v := "Sat, 28 Apr 2018 05:18:00 +0000"
	tm, err := parseTime(v)
	if err != nil {
		t.Error("TestParseTime: failed parse time, err: ", err)
	}

	if tm.Day() != 28 || tm.Month() != time.April || tm.Year() != 2018 ||
		tm.Hour() != 5 || tm.Minute() != 18 || tm.Second() != 0 {
		t.Error("TestParseTime: not expected time: ", tm)
	}

	// invalid
	v = "hoge"
	_, err = parseTime(v)
	if err == nil {
		t.Error("TestParseTime: expected error, but returned nil")
	}
}

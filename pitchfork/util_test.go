package pitchfork

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	// valid
	v := "Sat, 28 Apr 2018 05:18:00 +0000"
	t, err := parseTime(v)
	if err != nil {
		t.Error("TestParseTime: failed parse time, err: ", err)
	}

	if t.Day() != 28 || t.Month() != time.April || t.Yaer() != 2018 ||
		t.Hour() != 5 || t.Minute() != 18 || t.Second() != 0 {
		t.Error("TestParseTime: not expected time: ", t)
	}

	// invalid
	v = "hoge"
	_, err = parseTime(v)
	if err == nil {
		t.Error("TestParseTime: expected error, but returned nil")
	}
}

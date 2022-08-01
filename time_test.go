package ntime_test

import (
	"github.com/smartwalle/ntime"
	"testing"
	"time"
)

func TestTime_Greater(t *testing.T) {
	var testTbl = []struct {
		now    *ntime.Time
		dst    *ntime.Time
		expect bool
	}{
		{ntime.Now(), ntime.Now(), false},
		{ntime.Now(), ntime.Date(2018, time.May, 20, 0, 0, 0, 0, time.Local), true},
		{ntime.Now(), ntime.Now().Add(time.Second * 10), false},
		{ntime.Now(), nil, true},
	}

	for _, test := range testTbl {
		var actual = test.now.GreaterThan(test.dst)
		if actual != test.expect {
			t.Fatal(test.now, "比", test.dst, "大,", "期望:", test.expect, "实际:", actual)
		}
	}
}

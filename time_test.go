package time4go_test

import (
	"github.com/smartwalle/time4go"
	"testing"
	"time"
)

func TestTime_Greater(t *testing.T) {
	var testTbl = []struct {
		now    *time4go.Time
		dst    *time4go.Time
		expect bool
	}{
		{time4go.Now(), time4go.Now(), false},
		{time4go.Now(), time4go.Date(2018, time.May, 20, 0, 0, 0, 0, time.Local), true},
		{time4go.Now(), time4go.Now().Add(time.Second * 10), false},
		{time4go.Now(), nil, true},
	}

	for _, test := range testTbl {
		var actual = test.now.GreaterThan(test.dst)
		if actual != test.expect {
			t.Fatal(test.now, "比", test.dst, "大,", "期望:", test.expect, "实际:", actual)
		}
	}
}

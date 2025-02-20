package ntime_test

import (
	"github.com/smartwalle/ntime"
	"testing"
	"time"
)

func TestTime_After(t *testing.T) {
	var testTbl = []struct {
		now    ntime.Time
		dst    ntime.Time
		expect bool
	}{
		{now: ntime.Now(), dst: ntime.Now(), expect: false},
		{now: ntime.Now(), dst: ntime.Date(2018, time.May, 20, 0, 0, 0, 0, time.Local), expect: true},
		{now: ntime.Now(), dst: ntime.Now().Add(time.Second * 10), expect: false},
		{now: ntime.Now(), expect: true},
	}

	for _, test := range testTbl {
		var actual = test.now.After(test.dst)
		if actual != test.expect {
			t.Fatal(test.now, "比", test.dst, "大,", "期望:", test.expect, "实际:", actual)
		}
	}
}

func TestTime_Beginning(t *testing.T) {
	var now = ntime.Now()
	t.Log("-Minute", now.BeginningOfMinute(), now.EndOfMinute())
	t.Log("---Hour", now.BeginningOfHour(), now.EndOfHour())
	t.Log("----Day", now.BeginningOfDay(), now.EndOfDay())
	t.Log("---Week", now.BeginningOfWeek(), now.EndOfWeek())
	t.Log("--Month", now.BeginningOfMonth(), now.EndOfMonth())
	t.Log("Quarter", now.BeginningOfQuarter(), now.EndOfQuarter())
	t.Log("---Year", now.BeginningOfYear(), now.EndOfYear())
}

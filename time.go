package ntime

import (
	"time"
)

var (
	JSONFormatter Formatter = DefaultFormatter{time.RFC3339}
)

const (
	kDefaultLayout = "2006-01-02 15:04:05.999999999 -0700 MST"
)

type Time struct {
	time.Time
}

func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}

func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
}

func (t Time) Compare(u Time) int {
	return t.Time.Compare(u.Time)
}

func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}

func (t Time) Add(d time.Duration) Time {
	return Time{Time: t.Time.Add(d)}
}

func (t Time) Sub(u Time) time.Duration {
	var t1 = t.Time
	var t2 = u.Time
	t1 = t1.In(time.UTC)
	t2 = t2.In(time.UTC)
	return t1.Sub(t2)
}

func (t Time) AddDate(years int, months int, days int) Time {
	return Time{Time: t.Time.AddDate(years, months, days)}
}

func (t Time) UTC() Time {
	return Time{Time: t.Time.UTC()}
}

func (t Time) Local() Time {
	return Time{Time: t.Time.Local()}
}

func (t Time) In(loc *time.Location) Time {
	return Time{Time: t.Time.In(loc)}
}

func (t Time) Truncate(d time.Duration) Time {
	return Time{Time: t.Time.Truncate(d)}
}

func (t Time) Round(d time.Duration) Time {
	return Time{Time: t.Time.Round(d)}
}

func (t Time) Format(layout string) string {
	if layout == "" {
		layout = kDefaultLayout
	}
	return t.Time.Format(layout)
}

// UnixNanosecond 纳秒（ns）
func (t Time) UnixNanosecond() int64 {
	return t.UnixNano()
}

// UnixMicrosecond 微秒（µs）
func (t Time) UnixMicrosecond() int64 {
	return t.UnixNano() / 1e3
}

// UnixMillisecond 毫秒（ms）
func (t Time) UnixMillisecond() int64 {
	return t.UnixNano() / 1e6
}

// UnixSecond 秒（s）
func (t Time) UnixSecond() int64 {
	return t.Unix()
}

// Previous 获取当前日期的前一天（昨天）
func (t Time) Previous() Time {
	return Time{Time: t.Time.Add(time.Hour * -24)}
}

// Next 获取当前日期的后一天（明天）
func (t Time) Next() Time {
	return Time{Time: t.Time.Add(time.Hour * 24)}
}

// FirstDayOfWeek 获取当前日期所在周的第一天
func (t Time) FirstDayOfWeek() Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var w = nt.Weekday()
	var d = int(w - time.Sunday)
	return Date(t.Year(), t.Month(), t.Day()-d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// LastDayOfWeek 获取当前日期所在周的最后一天
func (t Time) LastDayOfWeek() Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	var w = nt.Weekday()
	var d = int(time.Saturday - w)
	return Date(t.Year(), t.Month(), t.Day()+d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// FirstDayOfMonth 获取当前日期所在月的第一天
func (t Time) FirstDayOfMonth() Time {
	return Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// LastDayOfMonth 获取当前日期所在月的最后一天
func (t Time) LastDayOfMonth() Time {
	return Date(t.Year(), t.Month(), NumberOfDaysInMonth(t.Year(), t.Month()), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// FirstDayOfYear 获取当前日期所在年的第一天
func (t Time) FirstDayOfYear() Time {
	return Date(t.Year(), time.January, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// LastDayOfYear 获取当前日期所在年的最后一天
func (t Time) LastDayOfYear() Time {
	return Date(t.Year(), time.December, NumberOfDaysInMonth(t.Year(), time.December), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// Start 获取当前日期的开始时间，如：2023-10-15 00:00:00 +0000 UTC
func (t Time) Start() Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// End 获取当前日期的结束时间，如：2023-10-15 23:59:59.999999999 +0000 UTC
func (t Time) End() Time {
	return Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// TrimSecond 将 Second 及以下的单位清除，只保留 Minute 及以上的信息
func (t Time) TrimSecond() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// TrimMinute 将 Minute 及以下的单位清除，只保留 Hour 及以上的信息
func (t Time) TrimMinute() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// TrimHour 将 Hour 及以下的单位清除，只保留 Day 及以上的信息
func (t Time) TrimHour() Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

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

// BeginningOfMinute 获取当前分钟的开始时间
func (t Time) BeginningOfMinute() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

func (t Time) EndOfMinute() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginningOfHour 获取当前小时的开始时间
func (t Time) BeginningOfHour() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 获取当前小时的结束时间
func (t Time) EndOfHour() Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginningOfDay 获取当前天的开始时间
func (t Time) BeginningOfDay() Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 获取当前天的结束时间
func (t Time) EndOfDay() Time {
	return Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginningOfWeek 获取当前日期所在周的开始时间
func (t Time) BeginningOfWeek() Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var w = nt.Weekday()
	var d = int(w - time.Sunday)
	return Date(t.Year(), t.Month(), t.Day()-d, 0, 0, 0, 0, t.Location())
}

// EndOfWeek 获取当前日期所在周的结束时间
func (t Time) EndOfWeek() Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	var w = nt.Weekday()
	var d = int(time.Saturday - w)
	return Date(t.Year(), t.Month(), t.Day()+d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginningOfMonth 获取当前日期所在月的开始时间
func (t Time) BeginningOfMonth() Time {
	return Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 获取当前日期所在月的结束时间
func (t Time) EndOfMonth() Time {
	return Date(t.Year(), t.Month(), NumberOfDaysInMonth(t.Year(), t.Month()), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginningOfYear 获取当前日期所在年的开始时间
func (t Time) BeginningOfYear() Time {
	return Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 获取当前日期所在年的结束时间
func (t Time) EndOfYear() Time {
	return Date(t.Year(), time.December, NumberOfDaysInMonth(t.Year(), time.December), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

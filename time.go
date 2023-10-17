package ntime

import (
	"database/sql/driver"
	"fmt"
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

func (t Time) MarshalBinary() ([]byte, error) {
	return t.Time.MarshalBinary()
}

func (t *Time) UnmarshalBinary(data []byte) error {
	return t.Time.UnmarshalBinary(data)
}

func (t Time) GobEncode() ([]byte, error) {
	return t.Time.GobEncode()
}

func (t *Time) GobDecode(data []byte) error {
	return t.Time.GobDecode(data)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return JSONFormatter.Format(t.Time)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = JSONFormatter.Parse(data)
	return err
}

func (t Time) MarshalText() ([]byte, error) {
	return t.Time.MarshalText()
}

func (t *Time) UnmarshalText(data []byte) error {
	return t.Time.UnmarshalText(data)
}

func (t *Time) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.Time.UTC(), nil
}

func (t *Time) Scan(value interface{}) (err error) {
	switch val := value.(type) {
	case time.Time:
		t.Time = val.UTC()
		return nil
	case *time.Time:
		t.Time = (*val).UTC()
		return nil
	case Time:
		t.Time = val.Time.UTC()
		return nil
	case *Time:
		t.Time = val.Time.UTC()
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("ntime: scanning unsupported type: %T", value)
	}
}

func (t Time) Format(layout string) string {
	if layout == "" {
		layout = kDefaultLayout
	}
	return t.Time.Format(layout)
}

func (t Time) GreaterThan(u Time) bool {
	return t.After(u)
}

func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}

func (t Time) LessThan(u Time) bool {
	return t.Before(u)
}

func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
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

func (t Time) AddDate(years int, months int, days int) Time {
	return Time{Time: t.Time.AddDate(years, months, days)}
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

// BeginningDateOfWeek 获取当前日期所在周的第一天
func (t Time) BeginningDateOfWeek() Time {
	return beginningDateOfWeek(t)
}

// EndDateOfWeek 获取当前日期所在周的最后一天
func (t Time) EndDateOfWeek() Time {
	return endDateOfWeek(t)
}

// BeginningDateOfMonth 获取当前日期所在月的第一天
func (t Time) BeginningDateOfMonth() Time {
	return Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

// EndDateOfMonth 获取当前日期所在月的最后一天
func (t Time) EndDateOfMonth() Time {
	return Date(t.Year(), t.Month(), NumberOfDaysInMonth(t.Year(), t.Month()), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func (t Time) BeginningTime() Time {
	return BeginningTime(t)
}

func (t Time) EndTime() Time {
	return EndTime(t)
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

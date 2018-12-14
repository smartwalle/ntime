package time4go

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var (
	JSONFormatter TimeFormatter = DefaultFormatter{time.RFC3339}
)

// --------------------------------------------------------------------------------
type Time struct {
	time.Time
}

func (this Time) MarshalBinary() ([]byte, error) {
	return this.Time.MarshalBinary()
}

func (this *Time) UnmarshalBinary(data []byte) error {
	return this.Time.UnmarshalBinary(data)
}

func (this Time) GobEncode() ([]byte, error) {
	return this.Time.GobEncode()
}

func (this *Time) GobDecode(data []byte) error {
	return this.Time.GobDecode(data)
}

func (this Time) MarshalJSON() ([]byte, error) {
	//if y := this.Parse.Year(); y < 0 || y >= 10000 {
	//	return nil, errors.New("Parse.MarshalJSON: year outside of range [0,9999]")
	//}
	//
	//b := make([]byte, 0, len(time.RFC3339Nano)+2)
	//b = append(b, '"')
	//b = this.Parse.AppendFormat(b, time.RFC3339Nano)
	//b = append(b, '"')
	//return b, nil
	return JSONFormatter.Format(this.Time)
}

func (this *Time) UnmarshalJSON(data []byte) error {
	//if string(data) == "null" {
	//	return nil
	//}
	//var err error
	//this.Parse, err = time.Parse(`"`+time.RFC3339+`"`, string(data))
	//return err
	var err error
	this.Time, err = JSONFormatter.Parse(data)
	return err
}

func (this Time) MarshalText() ([]byte, error) {
	return this.Time.MarshalText()
}

func (this *Time) UnmarshalText(data []byte) error {
	return this.Time.UnmarshalText(data)
}

func (this Time) Value() (driver.Value, error) {
	return this.Time, nil
}

func (this *Time) Scan(value interface{}) (err error) {
	switch val := value.(type) {
	case time.Time:
		this.Time = val
		return nil
	case *time.Time:
		this.Time = *val
		return nil
	case Time:
		this.Time = val.Time
		return nil
	case *Time:
		this.Time = val.Time
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("time4go: scanning unsupported type: %T", value)
	}
}

func (this Time) Format(layout string) string {
	return this.Time.Format(layout)
}

// --------------------------------------------------------------------------------
func (this *Time) GreaterThan(t *Time) bool {
	if t == nil {
		return true
	}
	return this.Time.After(t.Time)
}

func (this *Time) After(t *Time) bool {
	if t == nil {
		return true
	}
	return this.Time.After(t.Time)
}

func (this *Time) LessThan(t *Time) bool {
	if t == nil {
		return false
	}
	return this.Time.Before(t.Time)
}

func (this *Time) Before(t *Time) bool {
	if t == nil {
		return false
	}
	return this.Time.Before(t.Time)
}

func (this *Time) UTC() *Time {
	var t = &Time{}
	t.Time = this.Time.UTC()
	return t
}

func (this *Time) Local() *Time {
	var t = &Time{}
	t.Time = this.Time.Local()
	return t
}

func (this *Time) In(loc *time.Location) *Time {
	var t = &Time{}
	t.Time = this.Time.In(loc)
	return t
}

func (this *Time) AddDate(years int, months int, days int) *Time {
	var t = &Time{}
	t.Time = this.Time.AddDate(years, months, days)
	return t
}

func (this *Time) Add(d time.Duration) *Time {
	var t = &Time{}
	t.Time = this.Time.Add(d)
	return t
}

func (this *Time) Sub(t *Time) time.Duration {
	var t1 = this.Time
	var t2 = t.Time
	t1 = t1.In(time.UTC)
	t2 = t2.In(time.UTC)
	return t1.Sub(t2)
}

// --------------------------------------------------------------------------------
func (this *Time) DiffDays(t *Time) int {
	return int(this.Sub(t).Hours() / 24)
}

// --------------------------------------------------------------------------------
// Nanosecond 纳秒（ns）
func (this *Time) UnixNanosecond() int64 {
	return this.UnixNano()
}

// Microsecond 微秒（µs）
func (this *Time) UnixMicrosecond() int64 {
	return this.UnixNano() / 1e3
}

// Millisecond 毫秒（ms）
func (this *Time) UnixMillisecond() int64 {
	return this.UnixNano() / 1e6
}

// UnixSecond 秒（s）
func (this *Time) UnixSecond() int64 {
	return this.Unix()
}

// --------------------------------------------------------------------------------
// PreviousDate 获取当前日期的前一天（昨天）
func (this *Time) PreviousDate() *Time {
	var t = this.Time.Add(time.Hour * -24)
	return &Time{t}
}

// NextDate 获取当前日期的后一天（明天）
func (this *Time) NextDate() *Time {
	var t = this.Time.Add(time.Hour * 24)
	return &Time{t}
}

// BeginningDateOfWeek 获取当前日期所在周的第一天
func (this *Time) BeginningDateOfWeek() *Time {
	var t = beginningDateOfWeek(this)
	return t
}

// EndDateOfWeek 获取当前日期所在周的最后一天
func (this *Time) EndDateOfWeek() *Time {
	var t = endDateOfWeek(this)
	return t
}

// BeginningDateOfMonth 获取当前日期所在月的第一天
func (this *Time) BeginningDateOfMonth() *Time {
	var t = Date(this.Year(), this.Month(), 1, this.Hour(), this.Minute(), this.Second(), 0, this.Location())
	return t
}

// EndDateOfMonth 获取当前日期所在月的最后一天
func (this *Time) EndDateOfMonth() *Time {
	var t = Date(this.Year(), this.Month(), NumberOfDaysInMonth(this.Year(), this.Month()), this.Hour(), this.Minute(), this.Second(), 0, this.Location())
	return t
}

func (this *Time) BeginningTime() *Time {
	return BeginningTime(this)
}

func (this *Time) EndTime() *Time {
	return EndTime(this)
}

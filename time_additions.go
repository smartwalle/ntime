package ntime

import "time"

// Now 获取当前时间
func Now() Time {
	return Time{time.Now()}
}

func FromTime(t time.Time) Time {
	return Time{t}
}

func FromNanosecond(ns int64) Time {
	return Unix(0, ns)
}

func FromMicrosecond(ms int64) Time {
	return Unix(0, ms*int64(time.Microsecond))
}

func FromMillisecond(ms int64) Time {
	return Unix(0, ms*int64(time.Millisecond))
}

func FromSecond(s int64) Time {
	return Unix(s, 0)
}

// NumberOfDaysInMonth 获取指定月份有多少天
func NumberOfDaysInMonth(year int, month time.Month) (number int) {
	number = 30

	switch month {
	case time.January:
		number = 31
	case time.February:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			number = 29
		} else {
			number = 28
		}
	case time.March:
		number = 31
	case time.April:
		number = 30
	case time.May:
		number = 31
	case time.June:
		number = 30
	case time.July:
		number = 31
	case time.August:
		number = 31
	case time.September:
		number = 30
	case time.October:
		number = 31
	case time.November:
		number = 30
	case time.December:
		number = 31
	}
	return number
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

func Unix(sec int64, nsec int64) Time {
	return Time{time.Unix(sec, nsec)}
}

func UnixIn(sec int64, nsec int64, loc *time.Location) Time {
	return Time{time.Unix(sec, nsec).In(loc)}
}

// BeginningDateOfYear 获取指定年份的第一天
func BeginningDateOfYear(year int) Time {
	return Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

// EndDateOfYear 获取指定年份的最后一天
func EndDateOfYear(year int) Time {
	return Date(year, 12, 31, 0, 0, 0, 0, time.Local)
}

// BeginningDateOfMonth 获取指定月份的第一天
func BeginningDateOfMonth(year int, month time.Month) Time {
	return Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

// EndDateOfMonth 获取指定月份的最后一天
func EndDateOfMonth(year int, month time.Month) Time {
	return Date(year, month, NumberOfDaysInMonth(year, month), 0, 0, 0, 0, time.Local)
}

// BeginningDateOfWeek 获取指定日期所在周的第一天
func BeginningDateOfWeek(year int, month time.Month, day int) Time {
	var t = Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(w - time.Sunday)
	return Date(year, month, day-d, 0, 0, 0, 0, time.Local)
}

func beginningDateOfWeek(t Time) Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	var w = nt.Weekday()
	var d = int(w - time.Sunday)
	return Date(t.Year(), t.Month(), t.Day()-d, t.Hour(), t.Minute(), t.Second(), 0, time.Local)
}

// EndDateOfWeek 获取指定日期所有周的最后一天
func EndDateOfWeek(year int, month time.Month, day int) Time {
	var t = Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(time.Saturday - w)
	return Date(year, month, day+d, 0, 0, 0, 0, time.Local)
}

func endDateOfWeek(t Time) Time {
	var nt = Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	var w = nt.Weekday()
	var d = int(time.Saturday - w)
	return Date(t.Year(), t.Month(), t.Day()+d, t.Hour(), t.Minute(), t.Second(), 0, time.Local)
}

// BeginningTimeOfDay 获取指定日期的开始时间
func BeginningTimeOfDay(year int, month time.Month, day int) Time {
	return Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// EndTimeOfDay 获取指定日期的结束时间
func EndTimeOfDay(year int, month time.Month, day int) Time {
	return Date(year, month, day, 23, 59, 59, 0, time.Local)
}

// BeginningTime 获取指定日期的开始时间
func BeginningTime(t Time) Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndTime 获取指定日期的结束时间
func EndTime(t Time) Time {
	return Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func Parse(layout, value string) (Time, error) {
	if layout == "" {
		layout = kDefaultLayout
	}

	t, err := time.Parse(layout, value)
	return Time{t}, err
}

func MustParse(layout, value string) Time {
	t, err := Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseInLocation(layout, value string, loc *time.Location) (Time, error) {
	if layout == "" {
		layout = kDefaultLayout
	}

	t, err := time.ParseInLocation(layout, value, loc)
	return Time{t}, err
}

func MustParseInLocation(layout, value string, loc *time.Location) Time {
	t, err := ParseInLocation(layout, value, loc)
	if err != nil {
		panic(err)
	}
	return t
}

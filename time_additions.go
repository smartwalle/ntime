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

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

func Unix(sec int64, nsec int64) Time {
	return Time{time.Unix(sec, nsec)}
}

func UnixIn(sec int64, nsec int64, loc *time.Location) Time {
	return Time{time.Unix(sec, nsec).In(loc)}
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

// FirstDayOfWeek 获取指定日期所在周的第一天
func FirstDayOfWeek(year int, month time.Month, day int) Time {
	var t = Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(w - time.Sunday)
	return Date(year, month, day-d, 0, 0, 0, 0, time.Local)
}

// LastDayOfWeek 获取指定日期所有周的最后一天
func LastDayOfWeek(year int, month time.Month, day int) Time {
	var t = Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(time.Saturday - w)
	return Date(year, month, day+d, 0, 0, 0, 0, time.Local)
}

// FirstDayOfMonth 获取指定月份的第一天
func FirstDayOfMonth(year int, month time.Month) Time {
	return Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

// LastDayOfMonth 获取指定月份的最后一天
func LastDayOfMonth(year int, month time.Month) Time {
	return Date(year, month, NumberOfDaysInMonth(year, month), 0, 0, 0, 0, time.Local)
}

// FirstDayOfYear 获取指定年份的第一天
func FirstDayOfYear(year int) Time {
	return Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
}

// LastDayOfYear 获取指定年份的最后一天
func LastDayOfYear(year int) Time {
	return Date(year, time.December, NumberOfDaysInMonth(year, time.December), 0, 0, 0, 0, time.Local)
}

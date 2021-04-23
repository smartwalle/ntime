package time4go

import (
	"errors"
	"time"
)

type TimeFormatter interface {
	Format(t time.Time) ([]byte, error)
	Parse(data []byte) (time.Time, error)
}

type DefaultFormatter struct {
	Layout string
}

func (this DefaultFormatter) Format(t time.Time) ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("time4go.JSONFormatter: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(this.Layout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, this.Layout)
	b = append(b, '"')
	return b, nil
}

func (this DefaultFormatter) Parse(data []byte) (result time.Time, err error) {
	if string(data) == "null" {
		return result, errors.New("time4go.JSONFormatter: invalid time")
	}
	result, err = time.Parse(`"`+this.Layout+`"`, string(data))
	return result, err
}

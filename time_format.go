package ntime

import (
	"time"
)

type Formatter interface {
	Format(t time.Time) ([]byte, error)
	Parse(data []byte) (time.Time, error)
}

type DefaultFormatter struct {
	Layout string
}

func (formatter DefaultFormatter) Format(t time.Time) ([]byte, error) {
	b := make([]byte, 0, len(formatter.Layout)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, formatter.Layout)
	}
	b = append(b, '"')
	return b, nil
}

func (formatter DefaultFormatter) Parse(data []byte) (t time.Time, err error) {
	if string(data) == "null" {
		return t, nil
	}
	t, err = time.Parse(`"`+formatter.Layout+`"`, string(data))
	return t, err
}

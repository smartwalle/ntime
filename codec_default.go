package ntime

import (
	"time"
)

type DefaultCodec struct {
	Layout string
}

func (c DefaultCodec) JSONEncode(t time.Time) (b []byte, err error) {
	b = make([]byte, 0, len(c.Layout)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, c.Layout)
	}
	b = append(b, '"')
	return b, nil
}

func (c DefaultCodec) JOSNDecode(data []byte) (t time.Time, err error) {
	if string(data) == "null" {
		return t, nil
	}
	if len(data) > 1 {
		data = data[1 : len(data)-1]
	}
	t, err = time.Parse(c.Layout, string(data))
	return t, err
}

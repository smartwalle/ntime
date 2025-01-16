package ntime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var (
	codec Codec = DefaultCodec{Layout: time.RFC3339}
)

func UseCodec(c Codec) {
	if c == nil {
		c = DefaultCodec{Layout: time.RFC3339}
	}
	codec = c
}

func GetCodec() Codec {
	return codec
}

type Codec interface {
	JSONEncode(t time.Time) ([]byte, error)
	JOSNDecode(data []byte) (time.Time, error)
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
	return t.GetCodec().JSONEncode(t.Time)
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = t.GetCodec().JOSNDecode(data)
	return err
}

func (t Time) MarshalText() ([]byte, error) {
	return t.Time.MarshalText()
}

func (t *Time) UnmarshalText(data []byte) error {
	return t.Time.UnmarshalText(data)
}

func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
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

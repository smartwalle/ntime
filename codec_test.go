package ntime_test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/smartwalle/ntime"
	"testing"
)

func TestCodec_Gob(t *testing.T) {
	var now = ntime.Now()

	var t1 = ntime.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())

	var buf = &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(t1); err != nil {
		t.Fatal(err)
	}

	var t2 ntime.Time
	if err := gob.NewDecoder(buf).Decode(&t2); err != nil {
		t.Fatal(err)
	}

	if !t2.Equal(t1) {
		t.Fatal("时间不相等")
	}
}

func TestCodec_JSON(t *testing.T) {
	var now = ntime.Now()

	var t1 = ntime.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())

	var buf = &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(t1); err != nil {
		t.Fatal(err)
	}

	var t2 ntime.Time
	if err := json.NewDecoder(buf).Decode(&t2); err != nil {
		t.Fatal(err)
	}

	if !t2.Equal(t1) {
		t.Fatal("时间不相等")
	}
}

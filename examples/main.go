package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/ntime"
)

func main() {
	ntime.UseCodec(ntime.DefaultCodec{Layout: "2006-01-02 15:04:05"})

	var s = &Schedule{}
	s.Start = ntime.Now()
	s.End = s.Start.AddDate(0, 1, 0)

	sBytes, _ := json.Marshal(s)
	fmt.Println(string(sBytes))

	var ts = `{"start":"2019-12-10 09:59:21","end":"2019-12-10 09:59:22"}`
	var s2 *Schedule
	json.Unmarshal([]byte(ts), &s2)
	fmt.Println(s2.Start, s2.End)

	//var now = ntime.Now()
	//fmt.Println(now.BeginningOfMonth())
	//fmt.Println(now.EndOfMonth())
	//
	//var t = ntime.Date(2019, time.January, 1, 0, 0, 0, 0, ntime.UTC)

	//fmt.Println(t.Format("2006-01-02 15:04:05"))
}

type Schedule struct {
	Start ntime.Time `json:"start"`
	End   ntime.Time `json:"end"`
}

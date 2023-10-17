package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/ntime"
	"time"
)

func main() {
	ntime.JSONFormatter = ntime.DefaultFormatter{Layout: "2006-01-02 15:04:05"}

	var s = &Schedule{}
	s.Begin = ntime.Now()
	//s.End = s.Begin.AddDate(0, 1, 0)

	sBytes, _ := json.Marshal(s)
	fmt.Println(string(sBytes))

	var ts = `{"begin":"2019-11-10 09:59:21","end":"2019-12-10 09:59:21"}`
	var s2 *Schedule
	json.Unmarshal([]byte(ts), &s2)
	fmt.Println(s2.Begin, s2.End)

	//var now = ntime.Now()
	//fmt.Println(now.BeginningDateOfMonth())
	//fmt.Println(now.EndDateOfMonth())

	var t = ntime.Date(2019, time.January, 1, 0, 0, 0, 0, ntime.UTC)

	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

type Schedule struct {
	Begin ntime.Time `json:"begin"`
	End   ntime.Time `json:"end"`
}

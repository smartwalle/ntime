package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/time4go"
)

func main() {
	time4go.JSONFormatter = time4go.DefaultFormatter{Layout: "2006-01-02 15:04:05"}

	var s = &Schedule{}
	s.Begin = time4go.Now()
	s.End = s.Begin.AddDate(0, 1, 0)

	sBytes, _ := json.Marshal(s)
	fmt.Println(string(sBytes))

	var ts = `{"begin":"2019-11-10 09:59:21","end":"2019-12-10 09:59:21"}`
	var s2 *Schedule
	json.Unmarshal([]byte(ts), &s2)
	fmt.Println(s2.Begin, s2.End)

	//var now = time4go.Now()
	//fmt.Println(now.BeginningDateOfMonth())
	//fmt.Println(now.EndDateOfMonth())
}

type Schedule struct {
	Begin *time4go.Time `json:"begin"`
	End   *time4go.Time `json:"end"`
}

package main

import (
	"fmt"
	"github.com/smartwalle/time4go"
	"time"
)

func main() {
	var now = time4go.Now()
	fmt.Println(now.BeginningDateOfMonth())
	fmt.Println(now.EndDateOfMonth())

	var d = time4go.Date(2018, time.May, 20, 13, 14, 0, 0, time.Local)
	fmt.Println(d)
}

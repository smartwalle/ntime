package main

import (
	"fmt"
	"github.com/smartwalle/time4go/timewheel"
	"sync"
	"time"
)

func main() {
	var tw = timewheel.New(time.Millisecond, 10)
	tw.Run()

	var wg = &sync.WaitGroup{}

	wg.Add(1)
	var b = time.Now().UnixMilli()
	tw.AfterFunc(time.Millisecond*1, func() {
		fmt.Println(time.Now().UnixMilli()-b, "done")
		wg.Done()
	})

	wg.Add(1)
	tw.AfterFunc(time.Millisecond*1100, func() {
		fmt.Println(time.Now().UnixMilli()-b, "done")
		wg.Done()
	})

	wg.Add(1)
	tw.AfterFunc(time.Millisecond*2200, func() {
		fmt.Println(time.Now().UnixMilli()-b, "done")
		wg.Done()
	})

	wg.Wait()
}

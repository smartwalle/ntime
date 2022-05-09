package main

import (
	"fmt"
	"github.com/smartwalle/time4go/timewheel"
	"sync"
	"time"
)

func main() {
	var tw = timewheel.New(time.Second, 10)
	tw.Run()

	var wg = &sync.WaitGroup{}

	wg.Add(1)
	tw.AfterFunc(time.Second*1, func() {
		fmt.Println(time.Now().UnixMilli()/1000, "done")
		wg.Done()
	})

	wg.Add(1)
	tw.AfterFunc(time.Second*13, func() {
		fmt.Println(time.Now().UnixMilli()/1000, "done")
		wg.Done()
	})

	wg.Add(1)
	tw.AfterFunc(time.Second*23, func() {
		fmt.Println(time.Now().UnixMilli()/1000, "done")
		wg.Done()
	})

	wg.Wait()
}

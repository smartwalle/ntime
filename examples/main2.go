package main

import (
	"fmt"
	"github.com/smartwalle/time4go/internal/timewheel"
	"sync"
	"time"
)

func main() {
	var tw = timewheel.New(time.Millisecond, 12)
	tw.Run()

	var wg = &sync.WaitGroup{}

	wg.Add(1)
	tw.After(time.Second*1, func() {
		fmt.Println(time.Now().Unix(), "done")
		wg.Done()
	})

	wg.Add(1)
	tw.After(time.Second*3, func() {
		fmt.Println(time.Now().Unix(), "done")
		wg.Done()
	})

	wg.Wait()
}

package main

import (
	"fmt"
	"github.com/smartwalle/time4go/internal/timewheel"
	"time"
)

func main() {
	tw := timewheel.New(200*time.Millisecond, 100)

	tw.Run()
	//defer tw.Stop()

	count := 500000
	queue := make(chan bool, count)

	// loop 3
	for index := 0; index < 3; index++ {
		start := time.Now()
		for index := 0; index < count; index++ {
			tw.After(time.Second*2, func() {
				queue <- true
			})
		}
		fmt.Println("add timer cost: ", time.Since(start))

		start = time.Now()
		incr := 0
		for {
			if incr == count {
				fmt.Println("recv sig cost: ", time.Since(start))
				break
			}

			<-queue
			incr++
		}
	}
}

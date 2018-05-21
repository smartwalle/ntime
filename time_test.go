package time4go

import (
	"fmt"
	"testing"
	"time"
)

func TestBeginningTimeOfWeek(t *testing.T) {
	fmt.Println(BeginningDateOfWeek(2018, time.May, 2), EndDateOfWeek(2018, time.May, 2))
	fmt.Println(BeginningDateOfWeek(2018, time.May, 15), EndDateOfWeek(2018, time.May, 17))
	fmt.Println(BeginningDateOfWeek(2018, time.May, 29), EndDateOfWeek(2018, time.May, 30))

	fmt.Println(Now().BeginningDateOfWeek(), Now().EndDateOfWeek())
}

func TestTime_Greater(t *testing.T) {
	var t1 = Now()
	var t2 = Date(2018, time.May, 20, 0, 0, 0, 0, time.Local)
	var t3 = Date(3018, time.May, 20, 0, 0, 0, 0, time.Local)

	fmt.Println(t1.GreaterThan(t2))
	fmt.Println(t1.GreaterThan(t3))
	fmt.Println(t1.GreaterThan(nil))
}

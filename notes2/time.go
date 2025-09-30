package notes2

import (
	"fmt"
	"time"
)

// time.Nanosecond
// time.Microsecond
// time.Millisecond
// time.Second
// time.Minute
// time.Hour

func TimeSleep() {
	for i := range 10 {
		time.Sleep(time.Millisecond * 200)
		fmt.Println(i)
	}
}

func TimeParse() {
	transform, err := time.ParseDuration("100000s")
	fmt.Println(transform, err)
}

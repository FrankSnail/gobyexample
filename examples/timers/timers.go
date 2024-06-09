// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	timer := time.NewTimer(time.Second)
	var t time.Time
	t = <-timer.C
	p(t)
	timer.Reset(time.Second)
	t = <-timer.C
	p(t)

	ticker := time.NewTicker(time.Second)
	for range 3 {
		t = <-ticker.C
		p(t)
	}
}

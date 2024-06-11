package examples

import (
	"fmt"
	"time"
)

func Timers_() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// output
// Timer 1 fired
// Timer 2 stopped

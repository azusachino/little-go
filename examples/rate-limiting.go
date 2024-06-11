package examples

import (
	"fmt"
	"time"
)

func RateLimit_() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.NewTicker(200 * time.Millisecond)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	bustyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		bustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			bustyLimiter <- t
		}
	}()

	bustyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bustyRequests <- i
	}

	close(bustyRequests)

	for req := range bustyRequests {
		<-bustyLimiter
		fmt.Println("request", req, time.Now())
	}
}

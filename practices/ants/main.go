package main

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.000"

type Pool struct {
	running int
}

func (p *Pool) increaseRunning() {
	p.running++
	fmt.Println(time.Now().Format(timeFormat))
}
func (p *Pool) decreaseRunning() {
	p.running--
	fmt.Println(time.Now().Format(timeFormat))
}

func Count(p *Pool) int {
	return p.running
}

func main() {
	timer := time.NewTimer(time.Second)
	var p = &Pool{}
	fmt.Println(Count(p))
	for {
		select {
		case <-timer.C:
			goto done
		default:
			p.run()
		}
	}
done:
	fmt.Println(Count(p))

	time.Sleep(time.Minute)
	fmt.Println(Count(p))
}

func (p *Pool) run() {
	p.increaseRunning()

	go func() {
		defer func() {
			p.decreaseRunning()
		}()
	}()
}

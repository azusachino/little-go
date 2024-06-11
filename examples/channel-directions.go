package examples

import "fmt"

// 有方向的channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings 输入流
// pongs 输出流
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

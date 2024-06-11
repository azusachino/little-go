package examples

import "fmt"

func NonBlocking_() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	default:
		fmt.Println("no message")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send msg:", msg)
	default:
		fmt.Println("no msg sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

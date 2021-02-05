package examples

import "fmt"

func main() {

	// 普通的channel
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

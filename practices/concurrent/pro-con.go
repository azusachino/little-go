package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main0() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)
	// 休眠无法保证稳定的输出结果
	time.Sleep(5 * time.Second)
}

func main_() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// Ctrl+C exit
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v) \n", <-sign)
}

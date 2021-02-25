package main

import (
	"fmt"
	"github.com/little-go/practices/concurrent"
	"strings"
	"sync"
	"time"
)

func main() {
	p := concurrent.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("Hello world")
	p.Publish("hello golang")

	go func() {
		for msg := range all {
			fmt.Println("all：", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang：", msg)
		}
	}()

	getRandom()

	// 单通道传值
	ch := make(chan bool)
	go worker(ch)
	time.Sleep(time.Second)
	ch <- true

	// 但是通道的发送操作和接收操作是一一对应的，如果要停止多个Goroutine，那么可能需要创建同样数量的通道，这个代价太大了。
	// 其实我们可以通过close()关闭一个通道来实现广播的效果，所有从关闭通道接收的操作均会收到一个零值和一个可选的失败标志。
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(ch2)
	}
	time.Sleep(time.Second)
	close(ch2)

	// 当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，因为main线程并没有等待各个工作Goroutine退出工作完成的机制。
	// 结合WG进行改进
	var wg sync.WaitGroup
	ch3 := make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker2(&wg, ch3)
	}
	time.Sleep(time.Second)
	close(ch3)
	wg.Wait()

	time.Sleep(time.Second * 10)
}

// 当有多个通道均可操作时，select会随机选择一个通道。
// 基于该特性我们可以用select实现一个生成随机数序列的程序
func getRandom() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// 基于select和default实现goroutine退出
func worker(ch chan bool) {
	for {
		select {
		default:
			// 正常操作
			fmt.Println("Doing My Job~")
		case <-ch:
			// 退出操作
			return
		}
	}
}

func worker2(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("Hello")
		case <-ch:
			return
		}
	}
}

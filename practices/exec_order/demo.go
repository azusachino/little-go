package main

import "sync"

var (
	a    string
	done bool
)

func setUp() {
	a = "Hello World"
	done = true
}

// Go语言并不保证在main()函数中观测到的对done的写入操作发生在对字符串a的写入操作之后，因此程序很可能打印一个空字符串。
// 更糟糕的是，因为两个线程之间没有同步事件，setup线程对done的写入操作甚至无法被main线程看到，main()函数有可能陷入死循环中。
func main() {
	go setUp()
	for !done {
	}
	print(a)
}

// 利用同步原语给两个事件明确排序
func main_() {
	// 当<-done执行时，必然要求done <- 1也已经执行。
	done := make(chan bool, 1)
	go func() {
		println("Hello World")
		done <- true
	}()
	<-done
}

// 函数的第二个mu.Lock()必然在后台线程的mu.Unlock()之后发生（sync.Mutex保证）
func _main_() {
	var mu sync.Mutex
	mu.Lock()

	go func() {
		println("Hello World")
		mu.Unlock()
	}()

	mu.Lock()
	// ... other actions
	defer mu.Unlock()
}

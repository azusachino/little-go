package concurrent

import (
	"fmt"
	"sync"
	"time"
)

// 无法保证执行顺序
func main1() {
	// mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型。
	var mu sync.Mutex
	go func() {
		fmt.Println("Hello World")
		mu.Lock()
	}()
	mu.Unlock()
}

// 通过mutex
func main2() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("Hello World")
		mu.Unlock()
	}()
	// Mutex保证此次Lock必在Unlock之后执行
	mu.Lock()
}

// 通过无缓冲channel实现同步
func main3() {
	// 根据Go语言内存模型规范，对于从无缓存通道进行的接收，发生在对该通道进行的发送完成之前。
	// 因此，后台线程<-done接收操作完成之后，main线程的done<- 1发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。
	done := make(chan int)
	go func() {
		fmt.Println("Hello World By Chan")
		<-done
	}()
	done <- 1
}

//上面的代码虽然可以正确同步，但是对通道的缓存大小太敏感：如果通道有缓存，就无法保证main()函数退出之前后台线程能正常打印了。
//更好的做法是将通道的发送和接收方向调换一下，这样可以避免同步事件受通道缓存大小的影响：
func main4() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("Hello World By Chan")
		done <- 1
	}()
	<-done
}

// 使用sync.WaitGroup来等待一组事件
func main5() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello WaitGroup")
			wg.Done()
		}()
	}

	// 等待N个后台进程完成
	wg.Wait()
}

func demo() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}

package concurrent

import (
	"fmt"
	"time"
)

// 带缓存通道的发送和接收规则来实现最大并发阻塞
var limit = make(chan int, 3)

type test func()

func a() {
	var tests []test
	tests = append(tests, func() {

	})
	for _, w := range tests {
		w := w
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	// 当select()有多个分支时，会随机选择一个可用的通道分支，如果没有可用的通道分支，则选择default分支，否则会一直保持阻塞状态。
	select {}
}

func b() {
	in := make(chan int)
	// 基于select实现通道的超时判断
	select {
	case v := <-in:
		fmt.Println(v)
	case <-time.After(time.Second):
		return // 已超时
	}

	// 通过select的default分支实现非阻塞的通道发送或接收操作
	select {
	case v := <-in:
		fmt.Println(v)
	default:
		// 没有数据的情况
	}

	// 通过空的select组织main函数退出
	select {}
}

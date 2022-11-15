package leak

import (
	"fmt"
	"runtime"
	"time"
)

func GetData() {
	// 未初始化
	var ch chan struct{}
	go func() {
		<-ch
	}()
}

func _main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	GetData()
	time.Sleep(2 * time.Second)
}

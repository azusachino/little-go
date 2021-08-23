package main

import (
	"fmt"
	"time"
)

func main() {
	timeAfterNew2()
}

// prints 4 3 2 1 0
func forLoop() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// block
func block() {
	// block ends
	// main ends
	// defer runs
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}
	fmt.Println("main ends")
}

// timeAfter 调用 defer 关键字会立刻拷贝函数中引用的外部参数，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，而是在 defer 关键字调用时计算的，最终导致上述代码输出 0s。
func timeAfter() {
	// 0s
	startAt := time.Now()
	defer fmt.Println(time.Since(startAt))

	time.Sleep(time.Second)
}

func timeAfterNew() {
	// 1.0007259s
	// 1.001122s
	startAt := time.Now()

	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(startAt)

	defer func() {
		fmt.Println(time.Since(startAt))
	}()

	time.Sleep(time.Second)
}

func timeAfterNew2() {
	// 1.0065277s
	// 1.0066013s
	startAt := time.Now()

	defer func() {
		fmt.Println(time.Since(startAt))
	}()

	defer func(t *time.Time) {
		fmt.Println(time.Since(*t))
	}(&startAt)

	time.Sleep(time.Second)
}

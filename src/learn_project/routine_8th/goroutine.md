# goroutine

## 协程

`go func(){}()`

- 轻量级'线程'
- 非抢占式多任务处理, 由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上运行

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from Go Routine %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
```

Subroutines are special cases of more general program components, called coroutines.

### goroutine的定义

- 任何函数只需加上go就能送给调度器运行
- 不需要在定义时区分是否是异步函数
- 调度器在合适的点进行切换
- 使用 -race检测数据访问的冲突

### goroutine可能的切换点

- I/O, select
- channel
- 等待锁
- 函数调用 (有时)
- runtime.Gosched()

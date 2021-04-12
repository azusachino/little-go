# timer

```go
package runtime

type timer struct {
	pp puintptr

	when     int64
	period   int64
	f        func(interface{}, uintptr)
	arg      interface{}
	seq      uintptr
	nextwhen int64
	status   uint32
}
```

- pp：计时器所在的处理器 P 的指针地址。
- when：计时器被唤醒的时间。
- period：计时器再次被唤醒的时间（when+period）。
- f：回调函数，每次在计时器被唤醒时都会调用。
- arg：回调函数的参数，每次在计时器被唤醒时会将该参数项传入回调函数 f 中。
- seq：回调函数的参数，该参数仅在 netpoll 的应用场景下使用。
- nextwhen：当计时器状态为 timerModifiedXX 时，将会使用 nextwhen 的值设置到 where 字段上。
- status：计时器的当前状态值，计时器本身包含大量的枚举标识
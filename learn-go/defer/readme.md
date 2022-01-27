# defer

## 执行顺序

多个 defer 的执行顺序为 LIFO (Stack)

## defer 声明时会先计算确定参数的值

```go
func a() {
    i := 0
    defer fmt.Println(i) // print 0
    i++ // no effect
    return
}
```

## defer 可以修改*有名返回值*函数的返回值

For instance, if the deferred function is a function literal and the surrounding function has named result parameters
that are in scope within the literal, the deferred function may access and modify the result parameters before they are
returned.

```go
// return 42
func a() (res int) {
    defer func() {
        re *= 7
    }()
    return 6
}

// return 100
func b() int {
    i := 100
    defer func() {
        i++
    }()
    return i
}
```

## defer 的类型

### 堆上分配

在 Go 1.13 之前所有  `defer`  都是在堆上分配，该机制在编译时：

1. 在  `defer`  语句的位置插入  `runtime.deferproc`，被执行时，`defer`  调用会保存为一个  `runtime._defer`  结构体，存入 Goroutine 的`_defer`
   链表的最前面；
2. 在函数返回之前的位置插入`runtime.deferreturn`，被执行时，会从 Goroutine 的  `_defer`  链表中取出最前面的`runtime._defer`  并依次执行。

### 栈上分配

Go 1.13 版本新加入  `deferprocStack`  实现了在栈上分配  `defer`，相比堆上分配，栈上分配在函数返回后  `_defer`  便得到释放，省去了内存分配时产生的性能开销，只需适当维护  `_defer`
的链表即可。

1.13 版本中并不是所有`defer`都能够在栈上分配。循环中的`defer`，无论是显示的`for`循环，还是`goto`形成的隐式循环，都只能使用堆上分配，即使循环一次也是只能使用堆上分配：

```go
func A1() {
    for i := 0; i < 1; i++ {
        defer println(i)
    }
}

$ GOOS=linux GOARCH=amd64 go tool compile -S main.go
        ...
        0x004e 00078 (main.go:5)        CALL    runtime.deferproc(SB)
        ...
        0x005a 00090 (main.go:5)        CALL    runtime.deferreturn(SB)
        0x005f 00095 (main.go:5)        MOVQ    32(SP), BP
        0x0064 00100 (main.go:5)        ADDQ    $40, SP
        0x0068 00104 (main.go:5)        RET
```

## 结构体

```go
type _defer struct {
    siz     int32       //参数和结果的内存大小
    started bool
    heap    bool        //是否是堆上分配
    openDefer bool      // 是否经过开放编码的优化
    sp        uintptr   //栈指针
    pc        uintptr   // 调用方的程序计数器
    fn        *funcval  // 传入的函数
    _panic    *_panic
    link      *_defer   //defer链表
    fd   unsafe.Pointer
    varp uintptr
    framepc uintptr
}
```

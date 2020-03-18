# 异常相关

## defer调用

延迟（defer）语句  
当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回  
在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前需要关闭相应的资源，不然很容易造成资源泄露等问题

```go
package main
import "os"
func ReadWrite() bool {
    file, err := os.Create("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    return true
}
```

## 错误处理

```go
package main
import (
"fmt"
"os"
)
func read() {
    file, err := os.Open("a.txt")
    defer file.Close()
    if err != nil {
        if pathError, ok := err.(*os.PathError); !ok {
        	panic(pathError)
        } else{
            fmt.Print("water")
        }
    }
}
```

## panic

- 停止当前函数执行
- 一直向上返回, 执行每一层的defer
- 如果没有遇见recover, 程序退出

## recover

- 仅在defer调用中使用
- 获取panic的值
- 如果无法处理, 可重新panic

## 错误处理综合

- defer + panic + recover
- Type Assertion
- 函数式编程的应用
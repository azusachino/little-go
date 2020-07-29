# INTERFACE

## duck typing

描述事物的外部行为而非内部结构  
go属于结构化类型系统

使用者 -> 实现者
接口由使用者定义

```go
package main
type A interface {
    Get(url string) string
}
```
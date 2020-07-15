# Object

## 面向对象

go语言不支持多态和继承

```go
package main
type treeNode struct {
	value       int
	left, right *treeNode
}
func (node *treeNode) setValue(n int) {
	node.value = n
}
```

- 改变内容必须使用指针接收者  
- 结构过大考虑使用指针接收者

## wrapper (camelCase)

- 首字母大写 public
- 首字母小写 private

## package

- src下的目录 都是一个package
- 每个package下只能有一个main方法
- 定义在gopath下的src文件夹下的package才能引入

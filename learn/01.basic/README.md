# first-part

## 变量定义 (2020-03-12)

- var i int = 2 （类型在变量名之后，可省略，编译器自动推断）
- i := 2（声明变量语法糖）
- var i, j = 2, "22" (可以一次声明多个变量)
- var (i=1 j=2) 函数外声明 可以用括号一次声明多个变量

## 内建变量类型

- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr(指针)
- byte, rune(相当于char)
- float32, float64, complex64, complex128

## 强制类型转换

```go
package test
import ("fmt" 
        "math")
func triangle() {
	var a, b  = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}
```

## 常量的定义

const数值可以作为各种类型使用
```go
package main
import "fmt"
func consts() {
    const filename = "abc.txt" // 首字母大写 
    const a, b = 3, 4
    fmt.Println(filename ,a + b)
}
```

## 枚举类型

```go
package a
const (
cpp = 0 // 使用iota, 后续的枚举值可以省略 (自增)
_
java = 1
python = 2
golang = 3)

const (
b = 1 << (10 * iota) //自增值的种子
kb
mb
gb
)
```

## 条件语句

- for, if 后面的条件没有括号
- if条件里也可以定义变量 (作用域)
- 没有while
- switch不需要break, 也可以之间switch多个条件

## 函数

- 返回值类型写在最后
- 可返回多个值
- 函数可作为参数
- 没有默认参数, 可选参数

```go
package main
func m(a,b int) (c,d int) {
    return a +b, a-b
}
```

## 指针

```go
package main
func swap2(a, b *int) {
	*b, *a = *a, *b
}
```

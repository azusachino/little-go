# little-go

## 学习golang的repo

- 配置path => C:\Go\bin
- GoPath => C:\Users\az\Go

### 变量类型 条件语句

十进制转二进制  
对二取模, 依次排放  
13 -> 1 -> 01 -> 101 -> 1101

### 广度优先算法

### String 替换符号的规则 （string interpolation)

- %s	the uninterpreted bytes of the string or slice
- %q	a double-quoted string safely escaped with Go syntax
- %p	base 16 notation, with leading 0x

### base functions

- len(v Type) int => 获取长度
- cap(v Type) int => 获取容量
- make(t Type, size ...IntegerType) Type => 用于构建slice, map, or chan (only)
- new(Type) *Type => 相当于java中的new, 构建对象
- complex(r, i FloatType) ComplexType => 构建复数
- real(c complexType) FloatType => 将复数转换成实数
- close(c chan <- Type) => 关闭channel
- panic(v interface{}) => stops normal execution of the current goroutine
- recover() interface{} => allows a program to manage behavior of a panicking goroutine
- print(args ...Type) => 打印
- println(args ...Type) => 换行打印

### 切换proxy

`go env -w GOPROXY=https://goproxy.cn,direct`

### 知识点

struct最有用的特征之一是能够制定字段名映射

```go
type User struct {
  Id        int       `json:"id"`
  Name      string    `json:"name"`
  Bio       string    `json:"about,omitempty"`
  Active    bool      `json:"active"`
  Admin     bool      `json:"-"`
  CreatedAt time.Time `json:"created_at"`
}

```

强制转型

```go
func (dv DefaultValidator) Validate(val interface{}) (bool, error) {
	return val.(bool), nil // val.(type) 
}
```

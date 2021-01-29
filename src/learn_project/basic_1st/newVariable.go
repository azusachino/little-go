package basic_1st

import "fmt"

type Fruit struct {
	Name  string
	Price int8
}

var mk = "wa"

var (
	l1 = ""
	l2 = "aw"
)

func NamePrefix(s string) {
	if len(s) > 10 {
		fmt.Printf("your name length is over 10 \r\n")
	} else {
		fmt.Println(s)
	}
	fmt.Println(mk, l1, l2)
}

func Variable() {
	// 新建的变量都有初始值
	var i int       // 0
	var s string    // empty
	var a, b = 1, 2 // 省略类型，编译器自动推断
	var e, f, g = "1", 2, true
	k, l := "kk", "ll" // 和var定义变量一样（语法糖）只能在函数内部使用
	fmt.Println(s, i)
	fmt.Println(a + b)
	fmt.Printf(e, f, g)
	fmt.Printf(k + l)

	// 数组
	var balance [10]float32
	fmt.Print(balance)

	// 指针 一个指针变量指向了一个值的内存地址。
	var fruitPtr *Fruit
	fmt.Print(fruitPtr)

	// 定义指针变量。
	// 为指针变量赋值。
	// 访问指针变量中指向地址的值。
}

func ptr() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

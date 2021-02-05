package basic

import (
	"fmt"
	"io/ioutil"
)

func bounded(v int) bool {
	if v > 100 {
		return true
	} else if v < 0 {
		return false
	} else {
		return true
	}
}

func Aa() {
	const filename = "README.md"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func Aa2() {
	const filename = "README.md"
	// 作用域只在if内部
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func Eval(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case ">>":
		result = a >> b
	case "<<":
		result = a << b
	case "^":
		result = a ^ b
	default:
		return 0, fmt.Errorf("unsupported operator " + op)
		// panic("unsupported operator") // 中断程序
	}
	return result, nil
}

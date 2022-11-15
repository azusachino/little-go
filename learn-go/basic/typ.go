package basic

import (
	"fmt"
	"sync"
)

type Printer interface {
	print()
}

type HelloPrinter struct {
}

func (hp *HelloPrinter) print() {
	fmt.Println("hello")
}

// 获取当前interface的能力
type MyTools struct {
	Printer
	sync.Mutex
}

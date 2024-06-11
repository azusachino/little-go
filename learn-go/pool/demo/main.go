package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func Init() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new person")
			return new(Person)
		},
	}
}

/**
creating new person
Get Pool Object:  &{}
Get Pool Object： &{first}
creating new person # 第三次Get时，Pool中没有对象，调用New函数创建新的对象
Get Pool Object： &{}
*/
func main() {
	// runtime.GOMAXPROCS(0)
	person := pool.Get().(*Person)
	fmt.Println("Get Pool Object: ", person)
	person.Name = "first"
	pool.Put(person)

	fmt.Println("Get Pool Object: ", pool.Get().(*Person))
	fmt.Println("Get Pool Object: ", pool.Get().(*Person))
}

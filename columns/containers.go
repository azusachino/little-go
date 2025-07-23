package columns

import (
	"container/list"
	"container/ring"
	"fmt"
)

func Test_list() {
	lst := list.New()
	// wrap with element inside the function

	// before every operation, check initialization (?)
	lst.PushBack("")

	ring := ring.New(10)
	ring.Do(func(a any) {
		fmt.Printf("%v", a)
	})
}

func Test_map() {
	mp := map[string]int{
		"one": 1,
	}
	v, ok := mp["one"]
	if ok {
		fmt.Printf("%v", v)
	} 
}

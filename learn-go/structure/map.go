package structure

import (
	"fmt"
	"sync"
)

func M1() {
	m1 := map[string]string{
		"az": "bc",
	}
	var m2 map[int]int64       // m2 == nil
	m3 := make(map[string]int) // m3 == empty map

	fmt.Println(m1, m2, m3)

	// traversing map
	for _, v := range m3 {
		fmt.Println(v)
	}

	// get value from map
	az := m1["az"]
	fmt.Println(az)
	// if the key is not exist return empty
	if a, ok := m1["aw"]; ok {
		// return two value (check key exist)
		fmt.Println(a)
	} else {
		fmt.Println("key not exists")
	}

	mm := make(map[string]map[string]int)
	add(mm, "hello", "CN")

	mk := make(map[Key]int)
	mk[Key{
		Path:    "hello",
		Country: "CN",
	}] = 1

	c := Container{
		sync.RWMutex{},
		make(map[string]int),
	}

	c.Lock()
	c._data["hello"] = 1
	c.Unlock()
}

// deal with multi-level map
func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	// always check sub-level initialization
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

type Key struct {
	Path, Country string
}

type Container struct {
	sync.RWMutex
	_data map[string]int
}

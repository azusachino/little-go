package structure_2nd

import "fmt"

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
}

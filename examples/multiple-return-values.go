package examples

import "fmt"

func values() (int, int) {
	return 3, 6
}

func init() {
	a, b := values()
	fmt.Println(a, b)
	_, c := values()
	fmt.Println(c)
}

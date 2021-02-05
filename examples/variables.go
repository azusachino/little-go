package examples

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // default init value
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

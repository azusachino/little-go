package queue

import "fmt"

func main() {
	var q Queue
	q.Push("water")
	fmt.Println(q.Pop())
}

package exec_order

import (
	"fmt"
	"testing"
)

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func TestOrder(t *testing.T) {
	fmt.Println(a, b, c, d)
}

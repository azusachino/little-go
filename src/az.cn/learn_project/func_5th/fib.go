package func_5th

import (
	"fmt"
	"io"
	"strings"
)

func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1<<10 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d \n", next)
	return strings.NewReader(s).Read(p)
}

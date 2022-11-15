package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Ola() {
	c := 3 + 4i
	a := cmplx.Abs(c)
	fmt.Println(a)
}

func Ola2() {
	fmt.Println(cmplx.Exp(math.E), 1i*math.Pi)
}

func Triangle() {
	var a, b int = 3, 4
	var c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	ola()
	ola2()
}

func ola() {
	c := 3 + 4i
	a := cmplx.Abs(c)
	fmt.Println(a)
}

func ola2() {
	fmt.Println(cmplx.Exp(math.E), 1i*math.Pi)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

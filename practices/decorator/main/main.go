package main

import (
	"fmt"
	. "github.com/azusachino/golong/practices/decorator"
)

func main() {
	Decorator(Hello)("Hello World")
	sum1 := TimedSumFunc(Sum1)
	sum2 := TimedSumFunc(Sum2)
	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))

}

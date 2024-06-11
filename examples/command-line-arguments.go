package examples

import (
	"fmt"
	"os"
)

// $ ./command-line-arguments a b c d
// [./command-line-arguments a b c d]
// [a b c d]
// c
func CLIs() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

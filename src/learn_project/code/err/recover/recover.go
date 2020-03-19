package main

import "fmt"

func tryRecover() {
	// instant invoke the function
	defer func() {
		res := recover()
		if err, ok := res.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do with %v \n", res))
		}
	}()
	//b := 0
	//a := 5 / b
	//fmt.Println(a) // Error occurred runtime error: integer divide by zero
	panic(123)
}

func main() {
	tryRecover()
}

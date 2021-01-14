package decorator

import "fmt"

func Decorator(f func(string)) func(string) {
	return func(s string) {
		fmt.Println("Decorating...")
		f(s)
		fmt.Println("Finished...")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

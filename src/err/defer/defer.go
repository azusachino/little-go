package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	// 3 -> 2 -> 1
	// FILO
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic(i)
		}
	}
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName,
		os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s already exists!!\n", fileName)
		fmt.Println("Error: ", err.Error())

		if pathError, ok := err.(*os.PathError); !ok {
			panic(pathError)
		} else {
			fmt.Println(pathError.Op, pathError.Err, pathError.Path)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, "this is magic")
	}
}

func main() {
	// writeFile("magic.md")
	tryDefer2()
	writeFile("a")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(convert2Bin(5))    // 101
	fmt.Println(convert2Bin(13))   // 1101
	fmt.Println(convert2Bin(2315)) // 100100001011
	fmt.Println(convert2Bin(0))    // 0
	printFile("README.md")
}

func sum(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	return sum
}

func convert2Bin(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

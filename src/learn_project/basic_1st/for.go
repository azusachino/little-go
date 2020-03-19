package basic_1st

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Sum(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	return sum
}

func Convert2Bin(n int) string {
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

func PrintFile(filename string) {
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

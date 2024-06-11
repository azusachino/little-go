package examples

import (
	"fmt"
	"os"
)

func Exit_() {

	// Note that the ! from our program never got printed.
	defer fmt.Println("!")

	os.Exit(3)
}

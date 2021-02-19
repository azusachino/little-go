package examples

import (
	"fmt"
	"os"
)

func main() {

	// Note that the ! from our program never got printed.
	defer fmt.Println("!")

	os.Exit(3)
}

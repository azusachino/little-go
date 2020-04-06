# Go Channel

## basic semantic

```go
package main

import "fmt"

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

```

- channel
- buffered channel
- range
- Communication Sequential Process

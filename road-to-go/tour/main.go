package main

import (
	"github.com/azusachino/little-go/road-to-go/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute failed: %v", err)
	}
}

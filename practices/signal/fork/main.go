package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	path := "/path/to/executable"
	args := []string{"-graceful"}
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{}
	err := cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}

}

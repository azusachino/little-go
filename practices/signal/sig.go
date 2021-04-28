package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var cmd = exec.Command("ping")

func init() {
	cmd.Args = []string{"ping", "baidu.com"}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func main() {
	signalChan := make(chan os.Signal, 1)
	// 2 & 15
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		fmt.Printf("catch kubernetes signal: %v, stop cmd: %v", sig, cmd.Args)
		cmd.Process.Signal(os.Interrupt)
	}()

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}
}

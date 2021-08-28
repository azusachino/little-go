package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go", "Help")
	phpCmd.StringVar(&name, "n", "Php", "Help")

	args := flag.Args()

	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}

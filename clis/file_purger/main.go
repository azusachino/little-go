package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// clean specific files with suffix
func main() {
	rootDirectory := "." // Start from the current directory
	suffix := ".html"
	MoveFiles(rootDirectory, suffix)
}

func DeleteFiles(rootDirectory, suffix string) error {
	return filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == suffix {
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Error deleting file: %s\n", path)
			} else {
				fmt.Printf("Deleted file: %s\n", path)
			}
		}
		return nil
	})
}

func MoveFiles(rootDirectory, suffix string) error {
	return filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("current path %s", path)
		if !info.IsDir() && filepath.Ext(path) == suffix {
			newPath := filepath.Base(path) // Get just the filename
			err := os.Rename(path, newPath)
			if err != nil {
				return fmt.Errorf("error moving file %s: %v", path, err)
			}
			fmt.Printf("Moved file: %s to %s\n", path, newPath)
		}
		return nil
	})
}

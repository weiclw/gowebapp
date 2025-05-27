package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// genPackage creates a directory and writes "hello" into a file
func genPackage(browser, app string) {
	dir := filepath.Join("foo", "bar")

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	filePath := filepath.Join(dir, "f1")
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("hello"); err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	fmt.Printf("Successfully wrote to %s\n", filePath)
}

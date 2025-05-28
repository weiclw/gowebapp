package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getAllDirs() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home dir: %v\n", err)
		return []string{}
	}

	intermediates := []string{
		"Application",
	}

	app_dir := home
	for _, part := range intermediates {
		app_dir = filepath.Join(app_dir, part)
	}

	relative_list := []string{
		"Contents",
		"Contents/MacOS",
		"Contents/Resources",
	}

	ret := []string{}
	for _, p := range relative_list {
		ret = append(ret, filepath.Join(app_dir, p))
	}

	return ret
}

// genPackage creates a directory and writes "hello" into a file
func genPackage(browser, app string) {
	dir := filepath.Join("foo", "bar")

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	// Creating all directories in the wrapper package.
        all_dirs := getAllDirs()
	if len(all_dirs) == 0 {
		fmt.Printf("Fails to get directory list")
		return
	}

	for _, path := range all_dirs {
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Failed to create dir: %v\n", err)
			return
		}
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

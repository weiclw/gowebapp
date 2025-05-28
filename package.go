package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getPlist(app, icon string) string {
	header := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
`

	key_dict := map[string]string {
		"CFBundleName": app,
		"CFBundleDisplayName": app,
		"CFBundleIdentifier": "com.example." + app,
		"CFBundleVersion": "1.0",
		"CFBundleExecutable": "launcher.sh",
		"CFBundlePackageType": "APPL",
	}

	body := ""
	for k, v := range key_dict {
		body = body + "  <key>" + k + "</key>\n  <string>" + v + "</string>\n"
	}

	return header + "\n" +
		"<plist version=\"1.0\">\n" +
		"<dict>\n" +
		body +
		"</dict>\n" +
		"</plist>"
}

func getProfileRootDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home dir: %v\n", err)
		return ""
	}

	return filepath.Join(home, "profiles")
}

func getLauncher(browser string, app string, singleWindow bool) string {
	app_str := "https://www/" + app + ".com"
	if !singleWindow {
		app_str = "--app=\"" + app_str + "\""
	}

	profile_path := filepath.Join(getProfileRootDir(), app)

	return "#!/bin/bash\n\n" +
		"/Applications/" + app + ".app" + "/Contents/MacOS/" + app + " " +
		"--user-data-dir=" + profile_path + " " +
		app_str
}

func getAllDirs(app string) []string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home dir: %v\n", err)
		return []string{}
	}

	intermediates := []string{
		"Applications",
		app + ".app",
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

	ret = append(ret, filepath.Join(getProfileRootDir(), app))

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
    all_dirs := getAllDirs(app)
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

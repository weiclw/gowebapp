package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getPlist(app, icon string) string {
	header := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
`

	key_dict := map[string]string{
		"CFBundleName":        app,
		"CFBundleDisplayName": app,
		"CFBundleIdentifier":  "com.example." + app,
		"CFBundleVersion":     "1.0",
		"CFBundleExecutable":  "launcher.sh",
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
	app_str := "https://www." + app + ".com"
	if singleWindow {
		app_str = "--app=\"" + app_str + "\""
	}

	profile_path := filepath.Join(getProfileRootDir(), app)

	return "#!/bin/bash\n\n" +
		"\"/Applications/" + browser + ".app" + "/Contents/MacOS/" + browser + "\" " +
		"--user-data-dir=" + profile_path + " " +
		app_str
}

func getAppDir(app string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home dir: %v\n", err)
		return ""
	}

	intermediates := []string{
		"Applications",
		app + ".app",
	}

	app_dir := home
	for _, part := range intermediates {
		app_dir = filepath.Join(app_dir, part)
	}

	return app_dir
}

func getAllDirs(app string) []string {
	app_dir := getAppDir(app)
	if len(app_dir) == 0 {
		return []string{}
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

func addExecMode(filename string) {
	// Get the current file mode
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("failed to stat file: %v", err)
	}
	mode := info.Mode()

	// Add user execute permission (u+x)
	err = os.Chmod(filename, mode|0100)
	if err != nil {
		log.Fatalf("failed to chmod: %v", err)
	}
}

func genInfoplistFile(app_dir, app, icon string) {
	infoplist_path := filepath.Join(app_dir, "Contents")
	infoplist_path = filepath.Join(infoplist_path, "Info.plist")
	f, err := os.Create(infoplist_path)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer f.Close()

	infoplist_text := getPlist(app, icon)
	if _, err := f.WriteString(infoplist_text); err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	fmt.Printf("Successfully wrote to %s\n", infoplist_path)
}

func genLauncherFile(app_dir string, browser string, app string, singleWindow bool) {
	parts := []string{
		"Contents",
		"MacOS",
		"launcher.sh",
	}

	launcher_path := app_dir
	for _, part := range parts {
		launcher_path = filepath.Join(launcher_path, part)
	}

	f, err := os.Create(launcher_path)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}

	// Make sure that the file is closed at the end of the anonymous function.
	func(browser string, app string, singleWindow bool, f *os.File) {
		defer f.Close()

		launcher_text := getLauncher(browser, app, singleWindow)
		if _, err := f.WriteString(launcher_text); err != nil {
			fmt.Printf("Failed to write to file: %v\n", err)
			return
		}
	}(browser, app, singleWindow, f)

	addExecMode(launcher_path)

	fmt.Printf("Successfully wrote to %s\n", launcher_path)
}

// genPackage creates a directory and writes "hello" into a file
func genPackage(browser string, app string, icon string, singleWindow bool) {
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

	app_dir := getAppDir(app)
	if len(app_dir) == 0 {
		fmt.Printf("No app dir.")
		return
	}

	genInfoplistFile(app_dir, app, icon)
	genLauncherFile(app_dir, browser, app, singleWindow)
}

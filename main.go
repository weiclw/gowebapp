package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	browser := flag.String("browser", "", "Browser to use (e.g. chrome, brave)")
	app := flag.String("app", "", "App name or URL")
	profile := flag.String("profile", "", "Profile name to use")
	singleWindow := flag.Bool("single-window", false, "Prefer a single window")
	icon := flag.String("icon", "", "Icon file, can be png, etc")

	// Parse flags
	flag.Parse()

	// If no flags are passed, print usage info
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	// Output the parsed flag values
	fmt.Println("Launching with the following configuration:")
	fmt.Printf("  Browser:       %s\n", *browser)
	fmt.Printf("  App Name/URL:  %s\n", *app)
	fmt.Printf("  Profile Name:  %s\n", *profile)
	fmt.Printf("  Single Window: %t\n", *singleWindow)
	fmt.Printf("  Icon path:  %s\n", *icon)

	// Add your app logic here (e.g., construct a launch command or execute browser)
	genPackage(*browser, *app)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  --browser       Browser to use (e.g. chrome, brave)")
	fmt.Println("  --app           App name or URL to launch")
	fmt.Println("  --profile       Profile name or directory")
	fmt.Println("  --single-window If set, prefer a single window (true/false)")
	fmt.Println("  --icon          Icon file, can be png, etc")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  ./myapp --browser=chrome --app=https://chat.openai.com --profile=default --single-window")
}

package utils

import (
	"fmt"
	"os"
	"runtime"
)

func DetecOS() {
	switch so := runtime.GOOS; so {
	case "darwin", "linux":
		fmt.Println("🍎🐧")
	case "windows":
		fmt.Println("🪟")
	default:
		fmt.Println("no os detected")
		os.Exit(0)
	}
}

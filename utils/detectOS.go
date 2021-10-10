package utils

import (
	"fmt"
	"os"
	"runtime"
)

func DetecOS() {
	switch so := runtime.GOOS; so {
	case "darwin", "linux":
		fmt.Println("unix")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("no os detected")
		os.Exit(0)
	}
}

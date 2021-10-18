package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	switch runtime.GOOS {
	case "linux", "darwin":
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("No OS detected")
	}
}

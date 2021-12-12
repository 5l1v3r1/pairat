package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func KillProcess() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("killall", "ngrok")
		go func() {

			if err := cmd.Run(); err != nil {
				cmd = exec.Command("killall", "ngrok")
				if err := cmd.Run(); err != nil {
					log.Println(err)
				}
			}
		}()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			if err := cmd.Process.Kill(); err != nil {
				log.Println(err.Error())
			}
			os.Exit(0)
		}()
	case "windows":
		cmd := exec.Command("taskkill", "/F", "/IM", "ngrok.exe")
		if err := cmd.Run(); err != nil {
			log.Println(err)
		}
	default:
		fmt.Println("Platform not supported")
	}

}

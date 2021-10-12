package utils

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func KillProcess() {

	/*

		node (client)

	*/
	cmd := exec.Command("killall", "node")
	go func() {

		if err := cmd.Run(); err != nil {
			cmd = exec.Command("killall", "node")
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}
		}
	}()

	// stop the process when you type ctrl c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := cmd.Process.Kill(); err != nil {
			log.Println(err.Error())
		}
		os.Exit(0)
	}()
	/*

		Ngrok

	*/
	cmd2 := exec.Command("killall", "ngrok")
	go func() {

		if err := cmd2.Run(); err != nil {
			cmd = exec.Command("killall", "ngrok")
			if err := cmd2.Run(); err != nil {
				log.Println(err)
			}
		}
	}()

	// stop the process when you type ctrl c
	a := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-a
		if err := cmd.Process.Kill(); err != nil {
			log.Println(err.Error())
		}
		os.Exit(0)
	}()
}

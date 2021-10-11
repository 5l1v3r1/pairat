package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func Cli() {
	time.Sleep(1 * time.Second)
	// read/scan the file
	file, err := os.Open("src/utils/text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}

	/*

		This execute the client

	*/

	time.Sleep(2 * time.Second)
	fmt.Println("compilig client...")
	time.Sleep(1 * time.Second)
	cmd := exec.Command("npm", "start")
	go func() {

		if err := cmd.Run(); err != nil {
			cmd = exec.Command("npm", "start")
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

	fmt.Println("client on http://127.0.0.1:8000")
}

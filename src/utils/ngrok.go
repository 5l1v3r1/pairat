package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

var (
	detectNgrok     = regexp.MustCompile(`(https:)([/|.|\w|\s|-])*\.(?:io)`) // this is the regex for get the url
	cono        int = 0
)

func ExecuteNgrok() {

	/*

		Here is going execute ngrok

	*/
	cmd := exec.Command("ngrok", "http", "1323")
	go func() {

		if err := cmd.Run(); err != nil {
			cmd = exec.Command("ngrok", "http", "1323")
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

		Here this is why was doesen't working
		only wait 1 second to make the petition,
		so now wait 3 seconds.

	*/
	for i := 0; i < 3; i++ {
		fmt.Println("Searching the ngrok url...")
		time.Sleep(time.Second * 1)
	}

	res, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil && cono <= 10 {
		cono++

	} else if cono > 10 {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("url dont found")
		return
	}
	url := string(body)
	fmt.Println("Local client on: http://127.0.0.1:8000")
	fmt.Printf("\nPut this url in the remote cli client: \033[36m%s\n\n\033[0m", detectNgrok.FindString(url))

}

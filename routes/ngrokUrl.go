package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	detectNgrok = regexp.MustCompile(`(https:)([/|.|\w|\s|-])*\.(?:io)`) // this is the regex for get the url
)

func GetNgrokUrls() {

	res, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := detectNgrok.FindString(string(body))

	/*
		commands := url + "/commands"
		ip := url + "/ip"
		os := url + "/ip/os"
	*/

	fmt.Println("routes:")
	fmt.Println(url + "/commands")
	fmt.Println(url + "/ip")
	fmt.Println(url + "/ip/os")
}

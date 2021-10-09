package main

import (
	"github.com/ELPanaJose/pairat/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	// Home page
	e.GET("/", func(c echo.Context) error {
		c.File("view/index.html")
		return nil
	})
	e.POST("/commands", routes.UploadCommand)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

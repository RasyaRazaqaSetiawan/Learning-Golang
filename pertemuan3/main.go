package main

import (
	"net/http"
	"pertemuan3/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Assalaam")
	})

	routes.StudentRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}

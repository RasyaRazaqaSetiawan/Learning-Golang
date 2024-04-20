package routes

import (
	"pertemuan3/controller"

	"github.com/labstack/echo/v4"
)

func StudentRoute(e *echo.Echo) {
	e.POST("/student", controller.CreateStudent)
}

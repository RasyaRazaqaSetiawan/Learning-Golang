package controller

import (
	"log"
	"net/http"

	"pertemuan3/model"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func responseData(c echo.Context, status int, data interface{}, message string) error {
	return c.JSON(status, Response{
		Data:    data,
		Message: message,
	})
}

func CreateStudent(c echo.Context) (err error) {
	var s model.Student

	err = c.Bind(&s)
	if err != nil {
		log.Println("failed parse request")
		return
	}

	return responseData(c, http.StatusCreated, s, "Successfuly create student")
}

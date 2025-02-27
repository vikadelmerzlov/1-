package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

type RequestBody struct {
	Task string `json:"task"`
}

func PostHandler(c echo.Context) error {
	rBody := new(RequestBody)
	if err := c.Bind(rBody); err != nil {
		return c.JSON(http.StatusBadRequest, rBody)
	}
	return c.JSON(http.StatusOK, RequestBody{Task: task})
}

func main() {
	e := echo.New()
	e.POST("/api/hello", PostHandler)
	e.GET("/api/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Task!")
	})
	e.Start("localhost:8080")
}


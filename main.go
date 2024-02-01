package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Message{Message: "pong"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}

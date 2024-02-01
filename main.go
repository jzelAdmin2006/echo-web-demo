package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

type Result struct {
	Result int `json:"result"`
}

type Error struct {
	Error string `json:"error"`
}

func Plus() func(c echo.Context) error {
	return func(c echo.Context) error {
		a, err := strconv.Atoi(c.QueryParam("a"))
		b, err := strconv.Atoi(c.QueryParam("b"))
		if err != nil {
			return c.JSON(400, Error{
				Error: "invalid parameter",
			})
		}
		return c.JSON(http.StatusOK, Result{Result: a + b})
	}
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Message{Message: "pong"})
	})
	e.GET("/plus", Plus())
	e.Logger.Fatal(e.Start(":8080"))
}

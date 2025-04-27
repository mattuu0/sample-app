package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message" : "hello world",
		})
	})

	router.Logger.Fatal(router.Start(os.Getenv("BindAddr")))
}
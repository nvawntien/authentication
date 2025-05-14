package main

import (
	"myapp/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":8080"))
}

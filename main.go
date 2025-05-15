package main

import (
	"myapp/db"
	"myapp/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     1526,
		UserName: "vantien1526",
		Password: "Lpop3kiss",
		DbName:   "golang",
	}

	sql.Connect()

	defer sql.Close()

	e := echo.New()

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":8080"))
}

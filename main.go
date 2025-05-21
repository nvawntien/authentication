package main

import (
	"myapp/db"
	"myapp/handler"
	"myapp/repository/repo_impl"
	"myapp/router"

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

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":8080"))
}

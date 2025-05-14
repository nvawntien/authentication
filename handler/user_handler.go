package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "vantien",
		"email": "ngvantien1526@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string
		FullName string
		Age      int
	}

	user := User{
		Email:    "ngvantien1526@gmail.com",
		FullName: "Nguyen Van Tien",
		Age:      20,
	}

	return c.JSON(http.StatusOK, user)
}

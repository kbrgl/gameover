package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	digest, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user := User{
		Username:       username,
		PasswordDigest: string(digest),
		Email:          email,
	}
	tx := db.Create(&user)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "User created")
}

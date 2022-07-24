package main

import (
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	if username == "" || password == "" || email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing username, password, or email")
	}
	if len(username) < 2 || len(username) > 20 {
		return echo.NewHTTPError(http.StatusBadRequest, "username must be between 2 and 20 characters")
	}
	if len(password) < 5 {
		return echo.NewHTTPError(http.StatusBadRequest, "password must be at least 5 characters")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid email")
	}

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

	return c.NoContent(http.StatusCreated)
}

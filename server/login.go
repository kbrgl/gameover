package main

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing username or password")
	}

	var user User
	tx := db.First(&user, "username = ?", username)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username or password")
	}

	claims := jwt.StandardClaims{
		Subject: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, t)
}

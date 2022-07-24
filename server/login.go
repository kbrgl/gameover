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

	var user User
	tx := db.First(&user, "username = ?", username)
	if tx.Error != nil {
		return c.JSON(401, echo.Map{
			"message": "Unauthorized",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return c.JSON(401, echo.Map{
			"message": "Unauthorized",
		})
	}

	claims := jwt.StandardClaims{
		Subject: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

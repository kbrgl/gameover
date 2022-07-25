package main

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func friend(c echo.Context) error {
	friendUsername := c.FormValue("friend")
	if friendUsername == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing username")
	}

	var friend User
	err := db.Where("username = ?", friendUsername).First(&friend).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username")
	}

	var user User
	id := c.Get("user").(*jwt.Token).Claims.(*jwt.StandardClaims).Subject
	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user")
	}

	err = db.Model(&user).Association("Friends").Append(&friend)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to add friend")
	}

	return c.NoContent(http.StatusOK)
}

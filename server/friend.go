package main

import (
	"net/http"

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
	err = db.Model(&user).Association("Friends").Append(&friend)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to add friend")
	}
	return c.NoContent(201)
}

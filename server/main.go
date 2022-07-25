package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	secretKey []byte
)

func main() {
	secretKey = []byte(os.Getenv("SECRET_KEY"))

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&User{})

	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/friend", friend)

	e.Logger.Fatal(e.Start(":1323"))
}

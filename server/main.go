package main

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
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

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&User{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)
	e.POST("/register", register)
	config := middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: secretKey,
	}

	e.Use(middleware.JWTWithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}

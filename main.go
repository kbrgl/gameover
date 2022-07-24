package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kbrgl/gameover/cmd"
)

func main() {
	cmd.Execute()
}

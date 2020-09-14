package main

import (
	"Goletters/backend"
	"os"

	"github.com/joho/godotenv"
)

var server = backend.Server{}

func main() {
	godotenv.Load()
	Port := os.Getenv("PORT")
	server.Run(Port)
}

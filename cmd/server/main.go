package main

import (
	"os"

	"github.com/simonbuckner/go-web-server/internal/server"
)

func main() {

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	if len(port) == 0 {
		port = "1234"
	}
	server := server.New(host, port)
	server.Start()
}

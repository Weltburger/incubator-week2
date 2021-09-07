package main

import (
	"os"
	"simple-api/internal/server"
)

func main() {
	serv := server.NewServer()
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "1323"
	}
	serv.Router.Logger.Fatal(serv.Router.Start(":" + httpPort))
}

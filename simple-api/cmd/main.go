package main

import "simple-api/internal/server"

func main() {
	serv := server.NewServer()
	serv.Router.Logger.Fatal(serv.Router.Start(":1323"))
}

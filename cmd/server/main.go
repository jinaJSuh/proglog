package main

import (
	"server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fetal(srv.ListenAndServer())
}
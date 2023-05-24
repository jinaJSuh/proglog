package main

import (
	"github.com/JinaJSuh/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fetal(srv.ListenAndServer())
}
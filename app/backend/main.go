package main

import (
	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/server"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()

	router, err := server.NewBackendServer(cfg)
	if err != nil {
		panic(err)
	}

	log.Printf("server is listening at %s", ":10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}

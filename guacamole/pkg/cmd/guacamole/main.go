package main

import (
	"log"
	"net/http"

	"github.com/tatsuyayamauchi/golang-http-server-example/guacamole/pkg/api/router"
	"github.com/tatsuyayamauchi/golang-http-server-example/guacamole/pkg/cmd/guacamole/config"
)

func main() {
	cfg := config.Load()
	r := router.Initialize()

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

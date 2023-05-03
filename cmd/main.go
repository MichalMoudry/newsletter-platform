package main

import (
	"fmt"
	"log"
	"net/http"
	"newsletter-publishing/config"
	netchi "newsletter-publishing/transport"
)

func main() {
	cfg, cfgErr := config.ReadConfigFromFile("config.json")
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	handler := netchi.Initialize(cfg.Port)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

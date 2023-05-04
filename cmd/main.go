package main

import (
	"fmt"
	"log"
	"net/http"
	"newsletter-publishing/config"
	"newsletter-publishing/database"
	netchi "newsletter-publishing/transport"
)

func main() {
	cfg, cfgErr := config.ReadConfigFromFile("config.json")
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	if err := database.OpenDb(""); err != nil { //TODO: replace "" with cfg.ConnectionString
		log.Fatal(err)
	}

	handler := netchi.Initialize(cfg.Port)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

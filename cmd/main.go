package main

import (
	"fmt"
	"log"
	"net/http"
	"newsletter-platform/config"
	"newsletter-platform/database"
	netchi "newsletter-platform/transport"
	"newsletter-platform/transport/model"

	"github.com/go-chi/jwtauth/v5"
)

func main() {
	cfg, cfgErr := config.ReadConfigFromFile("../config.json")
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	if err := database.OpenDb(cfg.ConnectionString); err != nil {
		log.Fatal(err)
	}

	tokenAuth := jwtauth.New("HS512", []byte(cfg.SecurityString), nil)
	handler := netchi.Initialize(
		cfg.Port,
		tokenAuth,
		model.NewServiceCollection(tokenAuth),
	)

	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

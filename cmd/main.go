package main

import (
	"fmt"
	"log"
	"net/http"
	"newsletter-platform/config"
	"newsletter-platform/database"
	netchi "newsletter-platform/transport"
	"newsletter-platform/transport/model"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

func main() {
	cfg, cfgErr := config.ReadConfigFromFile("config.json")
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	if err := database.OpenDb(cfg.ConnectionString); err != nil {
		log.Fatal(err)
	}

	securityString := os.Getenv("SECURITY_STRING")
	if securityString == "" {
		securityString = "132456789-/ef-/eqw-f"
	}

	handler := netchi.Initialize(
		cfg.Port,
		jwtauth.New("HS512", []byte(securityString), nil),
		model.NewServiceCollection(),
	)

	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

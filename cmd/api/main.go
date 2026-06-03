package main

import (
	"log"
	"net/http"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/config"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/handler"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/repository"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/service"
	"github.com/joho/godotenv"
)

type App struct {
	Config config.Config
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file")
	}
	config.LoadConfig()

	mux := http.NewServeMux()

	fraudRepo := repository.NewFraudRepository()
	fraudSvc := service.NewFraudService(fraudRepo)
	fraudHandler := handler.NewFraudHandler(fraudSvc)

	readyHandler := handler.NewReadyHandler()

	mux.HandleFunc("POST /fraud-check", fraudHandler.FraudCheck)
	mux.HandleFunc("GET /ready", readyHandler.IsReady)

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}

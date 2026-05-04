package main

import (
	"log"
	"net/http"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ready", handler.Ready)

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}

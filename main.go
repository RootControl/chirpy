package main

import (
	"fmt"
	"net/http"

	"github.com/RootControl/chirpy/handlers"
)

func main() {
	mux := http.NewServeMux()
	api := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	link := "http://localhost" + api.Addr

	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("GET /metrics")
	fmt.Printf("Health Check: %s/health\n", link)

	api.ListenAndServe()
}

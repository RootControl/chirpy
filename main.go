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

	config := handlers.NewApiConfig()

	link := "http://localhost" + api.Addr

	mux.HandleFunc("GET /health", handlers.HealthCheck)
	fmt.Printf("Health Check: %s/health\n", link)

	mux.HandleFunc("GET /metrics", config.GetMetrics)
	fmt.Printf("Metrics: %s/metrics\n", link)

	mux.HandleFunc("POST /metrics/reset", config.ResetMetrics)
	fmt.Printf("Reset Metrics: %s/metrics/reset\n", link)

	err := api.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/RootControl/handlers"
)

func main() {
	mux := http.NewServeMux()
	api := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Listening on http://localhost:8080")

	handlers
	api.ListenAndServe()
}

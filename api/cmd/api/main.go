package main

import (
	"github.com/FilipeOliveira65/lookana/internal/handlers"
	"github.com/rs/cors"

	// "github.com/rs/cors"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-service", handlers.CreateService)
	mux.HandleFunc("/check-service", handlers.PingService)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

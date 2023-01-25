package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"github.com/user-management/script/router"
)

func main() {
	fmt.Println("run server")
	fmt.Println(os.Getenv("ENVIRONMENT"))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                                // All origins
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"}, // Allowing only get, just an example
		AllowedHeaders: []string{"access_token", "x-diug-ofni", "device_id", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	srv := &http.Server{
		Addr:    ":8901",
		Handler: c.Handler(router.Index),
	}

	srv.ListenAndServe()
}

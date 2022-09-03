package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.Load()
	r := router.Generate()
	/* disable cors */
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := cors.Handler(r)
	fmt.Println("Server is running on port:", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), handler))
}

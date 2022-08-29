package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()
	fmt.Println("Server is running on port:", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), r))
}

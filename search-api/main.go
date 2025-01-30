package main

import (
	"fmt"
	"log"
	"os"
	"search-api/initializers"
	"search-api/routes"

)

func init() {
	initializers.LoadEnv()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()
	log.Printf("Starting server on port %s...", port)
	r.Run(fmt.Sprintf(":%s", port))
}

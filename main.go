package main

import (
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatalln("PORT must be set")
	}

	router := chi.NewRouter()

}

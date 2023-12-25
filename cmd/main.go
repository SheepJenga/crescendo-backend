package main

import (
	"github.com/SheepJenga/crescendo/internal/server"
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

	svr, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Server listening on port", portStr)
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

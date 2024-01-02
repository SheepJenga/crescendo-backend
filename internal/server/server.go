package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

func newV1Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/healthz", handlerReadiness)
	router.Get("/err", handlerError)
	
	return router
}

func NewServer() (*http.Server, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatalln("PORT must be set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	v1Router := newV1Router()

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	return server, nil
}

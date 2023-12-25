package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

func newRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	return router
}

func NewServer() (*http.Server, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatalln("PORT must be set")
	}

	router := newRouter()

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	return server, nil

	//log.Println("Server listening on port", portStr)
	//err = server.ListenAndServe()
	//if err != nil {
	//	log.Fatalln(err)
	//}
}

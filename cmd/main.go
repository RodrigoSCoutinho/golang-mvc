package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	log.Printf("Starting server on port 8081")

	router.Get("/contact/", GetContact)
	router.Post("/contact", CreateContact)
	router.Put("/contact/{id}", UpdateContact)
	router.Delete("/contact/{id}", DeleteContact)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

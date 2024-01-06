package main

import (
	"golang-mvc/internal/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const port = ":8081"

func main() {
	router := chi.NewRouter()
	log.Printf("Starting server on port https://localhost%s", port)

	router.Get("/contact/", controller.GetContact)
	router.Post("/contact", controller.CreateContact)
	router.Put("/contact/{id}", controller.UpdateContact)
	router.Delete("/contact/{id}", controller.DeleteContact)

	log.Fatal(http.ListenAndServe(port, router))

}

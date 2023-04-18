package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("welcome chi\n"))

	})

	http.ListenAndServe(":8000", router)

}

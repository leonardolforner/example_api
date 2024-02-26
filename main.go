package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to the server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

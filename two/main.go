package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// chi router manages the handler
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// sets the route and handler to point to
	router.Get("/hello", basicHandler)

	// sets the server to listen to requests from the port and use the chi router as the handler
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// if there is a faulty request, print to console
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

// basicHandler responds to requests with a json response
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from the handler using go-chi!"))
}

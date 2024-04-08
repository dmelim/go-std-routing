package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Passing paremeters in the url
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("Receives request for item: " + id)))
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}

package main

import (
	"log"
	"net/http"

	"github.com/dmelim/go-std-routing/middleware"
)

func main() {
	router := http.NewServeMux()

	// Passing paremeters in the url
	router.HandleFunc("GET /item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("Received request for item: " + id)))
	})
	// Method based routing POST
	router.HandleFunc("POST /item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("We saved item:  " + id)))
	})
	// Method based routing DELETE
	router.HandleFunc("DELETE /item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("We deleted item:  " + id)))
	})
	// Can Handle host names instead of paths
	router.HandleFunc("hostname.com/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("We deleted item:  " + id)))
	})

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(v1),
	}
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}

package main

import (
	"log"
	"net/http"

	"github.com/dmelim/go-std-routing/middleware"
	"github.com/dmelim/go-std-routing/mythbeings"
)

func main() {
	router := http.NewServeMux()
	handler := &mythbeings.Handler{}

	// Passing paremeters in the url
	router.HandleFunc("GET /creature/{id}", handler.FindByID)
	// Method based routing POST
	router.HandleFunc("POST /creature", handler.Create)
	// Method based routing DELETE
	router.HandleFunc("DELETE /creature/{id}", handler.DeleteByID)

	// Can Handle host names instead of paths
	router.HandleFunc("hostname.com/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write(([]byte("We deleted item:  " + id)))
	})

	// Using another router to introduce subrouting
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(v1),
	}
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}

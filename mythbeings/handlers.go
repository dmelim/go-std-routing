package mythbeings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create a being")
	w.Write([]byte("Being created!"))
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling READ request - Method:", r.Method)

	being, exists := loadBeings()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(being)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling UPDATE request - Method:", r.Method)
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Received DELETE request for being")
	w.WriteHeader(http.StatusNotImplemented)
	being, exists := loadBeings()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	beingJSON, err := json.Marshal(being)
	if err != nil {
		// Handle error. For example, send an internal server error response.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling being into JSON"))
		return
	}

	// Construct the message string including the being JSON.
	message := fmt.Sprintf("Being: %s Deleted!", string(beingJSON))
	w.Write([]byte(message))
}

func (h *Handler) PatchByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

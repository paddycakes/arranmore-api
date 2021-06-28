package http

import (
	"fmt"
	"net/http"
)

import "github.com/gorilla/mux"

// Handler - stores pointer to metrics service
type Handler struct{
	Router *mux.Router
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all routes for application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
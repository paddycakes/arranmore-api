package http

import (
	"encoding/json"
	"fmt"
	"github.com/paddycakes/arranmore-api/internal/sensor"
	"net/http"
	"strconv"
)

import "github.com/gorilla/mux"

// Handler - stores pointer to metrics service
type Handler struct{
	Router *mux.Router
	Service *sensor.Service
}

// Response - an object to store responses from the api
type Response struct {
	Message string
	Error string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *sensor.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all routes for application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/sensor/{id}/metrics", h.GetSensorMetrics).Methods("GET")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

// GetSensorMetrics - retrieve metrics for sensor ID
func (h *Handler) GetSensorMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	vars := mux.Vars(r)
	id := vars["id"]	// Will this be id / clientId ?

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UInt from ID", err)
	}

	metrics, err := h.Service.GetMetrics(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Sensor Metrics for Client ID", err)
	}
	
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
package routes

import (
	"github.com/CagataySert/library-system/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
}

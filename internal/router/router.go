package router

import (
	"api-em-go/internal/handlers"

	"github.com/gorilla/mux"
)

// NewRouter cria e configura um novo roteador HTTP.
func NewRouter(userHandler *handlers.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Rotas de Health Check
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	// Rotas de Usuário
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	return router
}

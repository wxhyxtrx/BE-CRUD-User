package routes

import (
	"ptedi/handlers"
	"ptedi/pkg/connection"
	"ptedi/pkg/middleware"
	"ptedi/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(connection.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/checkauth", middleware.Auth(h.CheckAuth)).Methods("GET")
}

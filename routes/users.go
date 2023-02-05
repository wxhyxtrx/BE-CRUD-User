package routes

import (
	"ptedi/handlers"
	"ptedi/pkg/connection"

	"ptedi/pkg/middleware"
	"ptedi/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userReposetory := repositories.RepositoryUser(connection.DB)
	h := handlers.HandlerUser(userReposetory)

	r.HandleFunc("/users", h.FindAllUser).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}

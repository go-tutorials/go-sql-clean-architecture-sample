package http

import (
	"github.com/gorilla/mux"

	"go-sample/internal/user"
)

func UserRoutes(r *mux.Router, h user.UserPort) error {
	user := "/users"
	r.HandleFunc(user+"/search", h.Search).Methods("GET", "POST")
	r.HandleFunc(user+"/{id}", h.Load).Methods("GET")
	r.HandleFunc(user, h.Create).Methods("POST")
	r.HandleFunc(user+"/{id}", h.Update).Methods("PUT")
	r.HandleFunc(user+"/{id}", h.Patch).Methods("PATCH")
	r.HandleFunc(user+"/{id}", h.Delete).Methods("DELETE")

	return nil
}

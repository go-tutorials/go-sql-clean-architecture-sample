package http

import (
	"go-sample/internal/user"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router, h user.UserHandler) error {
	user := "/users"
	r.HandleFunc(user+"/search", h.Search).Methods("GET", "POST")
	r.HandleFunc(user+"/{id}", h.Load).Methods("GET")
	r.HandleFunc(user, h.Create).Methods("POST")
	r.HandleFunc(user+"/{id}", h.Update).Methods("PUT")
	r.HandleFunc(user+"/{id}", h.Patch).Methods("PATCH")
	r.HandleFunc(user+"/{id}", h.Delete).Methods("DELETE")

	return nil
}

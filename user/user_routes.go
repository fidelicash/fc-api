package user

import (
	"github.com/gorilla/mux"
)

// SetRoutes add routes from user
func SetRoutes(r *mux.Router) {
	r.HandleFunc("/transaction", AddTrsctn).Methods("POST")

}

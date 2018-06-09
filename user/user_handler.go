package user

import (
	"net/http"
)

// GetAll Users
func GetAll(w http.ResponseWriter, r *http.Request) {
	// db := db.Conn()
	// defer db.Close()
	// var users []User
	// db.Find(&users)
	// util.RespondWithJSON(w, http.StatusOK, users)
}

// FindByName find a User by name
func FindByName(w http.ResponseWriter, r *http.Request) {

	// util.RespondWithJSON(w, http.StatusOK, msg)

}

// FindByID find a User by ID
func FindByID(w http.ResponseWriter, r *http.Request) {

}

// Add a User
func Add(w http.ResponseWriter, r *http.Request) {

}

// DeleteByID find a User by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {

}

// UpdateByID find a User by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {

}

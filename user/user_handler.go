package user

import (
	"encoding/json"
	"net/http"

	"github.com/fidelicash/fc-api/util"
)

// Add a User
func AddTrsctn(w http.ResponseWriter, r *http.Request) {

	var trsctn Transaction
	var msg util.Message

	msg = util.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&trsctn); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	NewTransaction(trsctn)

	util.RespondWithJSON(w, http.StatusOK, "")
}

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

// DeleteByID find a User by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {

}

// UpdateByID find a User by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {

}

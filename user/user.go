package user

import (
	"fmt"

	"github.com/fidelicash/fc-api/firebase"
)

// User is a Human
type User struct {
	Name     string `json:"name"`
	FaceID   string `json:"age"`
	GoogleID string `json:"age"`
	CPF      string
	Balance  float32
}

func FindAll() {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	data, err := firebase.Get(client, "/")
	if err != nil {
		return
	}

	fmt.Println(data)
	return
}

// Users is a slice for User
// var Users []User

package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fidelicash/fc-api/firebase"
)

// User is a Human
type User struct {
	Name     string  `json:"nome,omitempty"`
	FaceID   string  `json:"IDFacebook,omitempty"`
	GoogleID string  `json:"IDGoogle,omitempty"`
	CPF      string  `json:"cpf,omitempty"`
	Balance  float32 `json:"saldo,omitempty"`
}

// History is a transaction
type Transaction struct {
	Origin string    `json:"origin,omitempty"`
	Target string    `json:"target,omitempty"`
	Value  float32   `json:"value,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}

func FindAll() {
	client, err := firebase.Conn()
	if err != nil {
		return
	}

	ctx := context.Background()
	ref := client.NewRef("/users")
	var users map[string]User
	if err := ref.Get(ctx, &users); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(users)
	return
}

func Find(id string) {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	ctx := context.Background()
	ref := client.NewRef("/users/" + id)
	var user User
	if err := ref.Get(ctx, &user); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(user)
	return
}

func Change() {

}

func Add() {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	ctx := context.Background()
	ref := client.NewRef("/")
	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*User{
		"alanisawesome": {
			Name:     "JC",
			FaceID:   "test1",
			GoogleID: "test1",
			CPF:      "46546545454",
			Balance:  1000,
		},
		"gracehop": {
			Name:     "JC2",
			FaceID:   "test2",
			GoogleID: "test2",
			CPF:      "46546545454",
			Balance:  900,
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
}

func Push() {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	ctx := context.Background()
	ref := client.NewRef("/")
	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*User{
		"alanisawesome": {
			Name:     "JC",
			FaceID:   "test12",
			GoogleID: "test1",
			CPF:      "46546545454",
			Balance:  1000,
		},
		"gracehop": {
			Name:     "JC2",
			FaceID:   "test2",
			GoogleID: "test2",
			CPF:      "46546545454",
			Balance:  900,
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
}

func UpdateSaldo(id string, saldo float32) {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	ctx := context.Background()
	ref := client.NewRef("/")
	usersRef := ref.Child("users")
	if err := usersRef.Update(ctx, map[string]interface{}{
		id + "/saldo": saldo,
	}); err != nil {
		log.Fatalln("Error updating children:", err)
	}
}

func AddTransaction(transaction Transaction) {
	client, err := firebase.Conn()
	if err != nil {
		return
	}
	ctx := context.Background()
	ref := client.NewRef("/history")
	transactionRef := ref.Child("")
	if _, err := transactionRef.Push(ctx, transaction); err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
}

func NewTransaction(transaction Transaction) {
	transaction.Date = time.Now()
	fmt.Println(transaction)
	UpdateSaldo("5gjhTAFCR9arcfSJuwthEzYikJF2", 200)
	Find("5gjhTAFCR9arcfSJuwthEzYikJF2")
	AddTransaction(transaction)
}

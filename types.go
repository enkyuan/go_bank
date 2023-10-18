package main

import (
	"math/rand"
	"time"
)

type createAccountRequest struct {
	FirstName 	string	`json:"firstName"`
	LastName 	string	`json:"lastName"`
}

type account struct {
	ID			int			`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Number		int64		`json:"number"`
	Balance 	int64		`json:"balance"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

func newAccount(firstName, lastName string) *account {
	return &account {
		ID: 		rand.Intn(10000),
		FirstName: 	firstName,
		LastName: 	lastName,
		Number: 	int64(rand.Intn(1000000)),
		CreatedAt: 	time.Now().UTC(),
	}
}
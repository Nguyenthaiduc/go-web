package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func Struct() {
	user := User{
		FirstName:   "Duc",
		LastName:    "Nguyen",
		PhoneNumber: "012345678",
		Age:         20,
	}
	log.Println(user.FirstName, user.LastName, "Phone Number :", user.PhoneNumber)
}

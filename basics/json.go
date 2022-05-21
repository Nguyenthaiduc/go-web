package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func JsonMain() {
	myJson := `
		[
			{
				"first_name" : "Clark",
				"last_name" : "Kent",
				"hair_color" : "black",
				"has_dog" : true
			},
			{
				"first_name" : "Elon",
				"last_name" : "Musk",
				"hair_color" : "brown",
				"has_dog" : false
			}
	    ]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error Unmarshal Json", err)
	}

	log.Printf("unmarshalled %v", unmarshalled)

	//write json a from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "John"
	m1.LastName = "Doe"
	m1.HairColor = "black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Elon"
	m2.LastName = "Musk"
	m2.HairColor = "red"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")

	if err != nil {
		log.Println("error marshalling")
	}

	fmt.Println(string(newJson))

}

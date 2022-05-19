package main

import "log"

type User2 struct {
	FirstName string
	LastName string
}

func Map() {

	me := User2 {
		FirstName: "Terra",
		LastName: "Luna",
	}

	myMapStruct := make(map[string]User2)

	myMapStruct["user"] = me

	log.Println(myMapStruct["user"].FirstName)


	//declare map
	myMap := make(map[string]string) //key and value

	myMap["dog"] = "Shiba"

	myMap["cat"] = "Luna"
	
	myMap["dog"] = "Shjba"

	log.Println(myMap["dog"])
	//Array

	var mySlice []string

	mySlice = append(mySlice,"Terra")
	mySlice = append(mySlice,"Luna")

	log.Println(mySlice)

	numbers := []int{1,2,3,4,5}

	log.Println(numbers)

	log.Println(numbers[0])
}

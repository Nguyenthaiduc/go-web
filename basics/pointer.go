package main

import "log"

func Pointer() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func callback myString is set to",myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to",s)
	newValue := "Red"
	*s = newValue
}

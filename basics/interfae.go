package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Age  int
}

type Gorilla struct {
	Name   string
	Age    int
	Weight int
}

func Interface() {
	dog := Dog{
		Name: "Shiba",
		Age:  5,
	}

	printInfo(dog)
	//
	gorilla := Gorilla{
		Name: "Gorilla",
		Age: 9,
		Weight: 70,
	}
	printInfo(gorilla)

}

func (d Dog) Says() string {
	return "I am Dog"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (d Gorilla) Says() string{
	return "I am Gorilla"
}

func (d Gorilla) NumberOfLegs() int{
	return 2
}


func printInfo(a Animal) {
	log.Println("The Animal says", a.Says(), "Number Leg :", a.NumberOfLegs())
}

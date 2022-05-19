package main

import "log"

func Loop() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	//arrayy
	mySlice := []string{"cat", "dog", "hourse", "tiger", "lion"}

	//index,value
	for _, x := range mySlice {
		log.Println(x)
	}
	//map

	myMap := make(map[string]string)

	myMap["1"] = "One"
	myMap["2"] = "Two"
	myMap["3"] = "Three"
	myMap["4"] = "Four"

	for i,x := range myMap {
		log.Println(i,x)
	}
}

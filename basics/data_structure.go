package main

import "log"

func DataStructure() {
	//example 1
	var isTrue bool
	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	}else {
		log.Println("isTrue is ",isTrue)
	}
	//example 2
	myNum := 100
	
	if myNum > 99 && !isTrue {
		log.Println("number greater than 99 and isTrue set to true")
	}else if myNum < 100 && isTrue {
		log.Println("1")
	}else if myNum == 101 || isTrue {
		log.Println("2")
	}else if myNum > 1000 && isTrue ==false {
		log.Println("3")
	}

	//Switch

	myVar := "cat"

	switch myVar {
	case "cat" :
		log.Println("Cats is set cat")

	case "dog" :
		log.Println("Cat is set Dog")
	case "fish":
		log.Println("Cat is set to Fish")

	default :
		log.Println("Cat is no something !")
	}
 }

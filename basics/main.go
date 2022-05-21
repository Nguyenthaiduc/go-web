package main

import "log"

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay = saySomething("Hello")

	saySomethingElse = saySomething("Goodbye")

	log.Println(whatToSay)
	log.Println(saySomethingElse)
	log.Println(saySomething("Finally"))

	i = 7
	i = 8
	log.Println(i)

	Pointer()
	// Struct_Function()
	Map()
	Interface()
	Channel()
	JsonMain()
	Test()

	
}



func saySomething(s string) string {
	return s
}

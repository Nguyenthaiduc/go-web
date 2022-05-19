package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about the page 2+2 is %d", sum))
}

func addValue(x, y int) int {
	return x + y
}

//main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting apllication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

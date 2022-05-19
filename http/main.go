package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0,10.0)
	if err != nil {
		fmt.Fprintf(w,"Cannot to divide by 0")
		return
	}

	fmt.Fprint(w,fmt.Sprintf("%f divided by %f is %f",100.0,10.0,f))
}

func divideValue(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("cannot divide zero")
		return 0,err
	}
	result := x / y
	return result, nil
}

//main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", About)

	fmt.Println(fmt.Sprintf("Starting apllication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

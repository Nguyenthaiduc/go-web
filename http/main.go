package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var portNumber = ":8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Parsing Template", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	

	fmt.Println(fmt.Sprintf("Starting apllication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

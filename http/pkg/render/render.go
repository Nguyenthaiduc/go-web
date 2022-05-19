package render

import (
	"fmt"
	"net/http"
	"html/template"
)

//RenderTemplate render template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Parsing Template", err)
		return
	}
}

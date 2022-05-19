package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nguyenthaiduc/go-web/pkg/config"
	"github.com/Nguyenthaiduc/go-web/pkg/handlers"
	"github.com/Nguyenthaiduc/go-web/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	//get the template cache from the appconfig

	tc,err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	
	
	repo := handlers.NewRepo((&app))
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting apllication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

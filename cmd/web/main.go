package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nguyenthaiduc/go-web/pkg/config"
	"github.com/Nguyenthaiduc/go-web/pkg/handlers"
	"github.com/Nguyenthaiduc/go-web/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //Coookie lien tuc
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//get the template cache from the appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo((&app))
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting apllication on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"

	"github.com/r00ti/simpleWebApp/config"
	"github.com/r00ti/simpleWebApp/pkg/handlers"
	"github.com/r00ti/simpleWebApp/pkg/render"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create tempalte cache!")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	err = http.ListenAndServe("localhost:8080", nil)
}

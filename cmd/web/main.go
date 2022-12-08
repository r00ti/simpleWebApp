package main

import (
	"log"
	"net/http"

	"github.com/r00ti/simpleWebApp/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

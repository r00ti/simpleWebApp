package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is at the about page 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divice", Divide)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

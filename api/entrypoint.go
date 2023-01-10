package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Hi</h1>
	`)
}

func Sub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Sub</h1>
	`)
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/sub", Sub)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

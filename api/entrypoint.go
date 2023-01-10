package handler

import (
	"fmt"
	"log"
	"net/http"
)

func Sur(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Sur</h1>
	`)
	fmt.Fprintf(w, `
	<h1>Surrr</h1>
	`)
	http.HandleFunc("/sup", Sup)
}

func Sup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Sup</h1>
	`)
}

func Sub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Sub</h1>
	`)
	http.HandleFunc("/sup", Sup)
}

func Super(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Super</h1>
	`)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/super", Super)
	http.HandleFunc("/sub", Sub)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {
// 	http.HandleFunc("/", Handler)
// 	http.HandleFunc("/sub", Sub)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

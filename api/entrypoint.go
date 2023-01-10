package handler

import (
	"fmt"
	"net/http"
)

func Sur(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<div>
	<h1>Sur</h1>
	</div>
	`)
	fmt.Fprintf(w, `
	<div>
	<h1>Surrr</h1>
	</div>
	`)
}

// func Sup(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 	<h1>Sup</h1>
// 	`)
// }

// func Sub(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 	<h1>Sub</h1>
// 	`)
// 	http.HandleFunc("/sup", Sup)
// }

// func Super(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 	<h1>Super</h1>
// 	`)
// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	http.HandleFunc("/super", Super)
// 	http.HandleFunc("/sub", Sub)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

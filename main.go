package main

import (
	"html/template"
	"net/http"

	"alura-store/db/postgres/repositories"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := repositories.FindAll()
	temp.ExecuteTemplate(w, "Index", products)
}

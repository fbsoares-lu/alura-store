package controllers

import (
	"alura-store/db/postgres/repositories"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := repositories.FindAll()
	temp.ExecuteTemplate(w, "Index", products)
}

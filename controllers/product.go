package controllers

import (
	"alura-store/db/postgres/repositories"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := repositories.FindAll()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFormatted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro to convert price", err)
		}

		quantityFormatted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro to convert quantity", err)
		}

		repositories.Insert(name, description, priceFormatted, quantityFormatted)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	productIdFormatted, err := strconv.Atoi(productId)

	if err != nil {
		log.Println("Error when format the product id", err)
	}

	repositories.Delete(productIdFormatted)
	http.Redirect(w, r, "/", 301)
}

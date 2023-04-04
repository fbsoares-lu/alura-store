package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func connection() *sql.DB {
	var user string = os.Getenv("POSTGRES_USER")
	var dbname string = os.Getenv("POSTGRES_DATABASE")
	var password string = os.Getenv("POSTGRES_PASSWORD")

	connVariables := "user=" + user + " dbname=" + dbname + " password=" + password + " host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connVariables)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connection()

	selectProducts, err := db.Query("select * from alura.products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}

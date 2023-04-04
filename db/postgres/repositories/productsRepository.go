package repositories

import (
	"alura-store/db"
	"alura-store/models"
)

func FindAll() []models.Product {
	db := db.Connection()

	selectProducts, err := db.Query("select * from alura.products")

	if err != nil {
		panic(err.Error())
	}

	p := models.Product{}
	products := []models.Product{}

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
	defer db.Close()
	return products
}

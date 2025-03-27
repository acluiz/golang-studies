package models

import "golang-studies/store-manager/db"


type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.ConnectDatabase()

	query := "select * from products"
	productsQuery, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	product := Product{}
	products := []Product{}

	for productsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsQuery.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()

	return products
}
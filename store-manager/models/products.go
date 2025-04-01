package models

import (
	"golang-studies/store-manager/db"
)


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

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()

	return products
}

func AddProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insert, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if (err != nil) {
		panic(err)
	}

	insert.Exec(name, description, price, quantity)
	
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	delete, err := db.Prepare("delete from products where id=$1")

	if (err != nil) {
		panic(err)
	}

	delete.Exec(id)

	defer db.Close()
}
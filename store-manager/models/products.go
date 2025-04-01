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

func GetProduct(id string) Product {
	db := db.ConnectDatabase()

	query := "select * from products where id=$1"
	getProduct, err := db.Query(query, id)

	if err != nil {
		panic(err)
	}

	product := Product{}

	for getProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64
	
		err :=	getProduct.Scan(&id, &name, &description, &price, &quantity)
			
		if err != nil {
			panic(err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()

	return product
}

func GetProducts() []Product {
	db := db.ConnectDatabase()

	query := "select * from products order by id asc"
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

	query, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if (err != nil) {
		panic(err)
	}

	query.Exec(name, description, price, quantity)
	
	defer db.Close()
}

func UpdateProduct(id, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	query, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err)
	}

	query.Exec(name, description, price, quantity, id)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	query, err := db.Prepare("delete from products where id=$1")

	if (err != nil) {
		panic(err)
	}

	query.Exec(id)

	defer db.Close()
}
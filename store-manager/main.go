package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

type Product struct {
	Id					int
	Name 				string
	Description string
	Price 			float64
	Quantity 		int
}

func index(w http.ResponseWriter, r *http.Request) {
	db :=	connectDatabase()

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

	temp.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}

func connectDatabase() *sql.DB {
	conn := ""
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}


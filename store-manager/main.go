package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

type Product struct {
	Name 				string
	Description string
	Price 			float64
	Quantity 		int
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta street 00K", Description: "Camiseta street oversize - Coleção 00K", Price: 69.9, Quantity: 22},
		{Name: "Tênis street 00K", Description: "Tênis street all black - Coleção 00K", Price: 289.9, Quantity: 14},
	}

	temp.ExecuteTemplate(w, "Index", products)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}


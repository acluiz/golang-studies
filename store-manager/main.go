package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"golang-studies/store-manager/models"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}


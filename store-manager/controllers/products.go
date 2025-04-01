package controllers

import (
	"fmt"
	m "golang-studies/store-manager/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := m.GetProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	product := m.GetProduct(id)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		
		if err != nil {
			fmt.Println("Falha ao converter preço")
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
	
		if err != nil {
			fmt.Println("Falha ao converter quantidade")
		}

		m.AddProduct(name, description, price, quantity)

		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		
		if err != nil {
			fmt.Println("Falha ao converter preço")
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
	
		if err != nil {
			fmt.Println("Falha ao converter quantidade")
		}

		m.UpdateProduct(id, name, description, price, quantity)
		
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	m.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}
package controllers

import (
	m "golang-studies/store-manager/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := m.GetProducts()

	temp.ExecuteTemplate(w, "Index", products)
}
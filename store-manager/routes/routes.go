package routes

import (
	c "golang-studies/store-manager/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/novo-produto", c.NewProduct)
	
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/delete", c.Delete)
}
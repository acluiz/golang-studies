package routes

import (
	c "golang-studies/store-manager/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/novo-produto", c.NewProduct)
	http.HandleFunc("/editar-produto", c.Edit)
	
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/update", c.Update)
	http.HandleFunc("/delete", c.Delete)
}
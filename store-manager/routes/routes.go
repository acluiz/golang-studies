package routes

import (
	c "golang-studies/store-manager/controllers"
	"net/http"
)


func LoadRoutes() {
	http.HandleFunc("/", c.Index)
}
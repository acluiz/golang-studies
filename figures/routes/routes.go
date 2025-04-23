package routes

import (
	"log"
	"net/http"

	c "golang-studies/figures/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", c.Home)
	http.HandleFunc("/figures", c.GetFigures)
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}
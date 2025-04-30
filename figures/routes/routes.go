package routes

import (
	"log"
	"net/http"

	c "golang-studies/figures/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", c.Home)
	
	r.HandleFunc("/api/figures", c.GetFigures).Methods("Get")
	r.HandleFunc("/api/figures", c.CreateFigure).Methods("Post")
	
	r.HandleFunc("/api/figures/{id}", c.GetFigure).Methods("Get")
	r.HandleFunc("/api/figures/{id}", c.UpdateFigure).Methods("Put")
	r.HandleFunc("/api/figures/{id}", c.DeleteFigure).Methods("Delete")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}
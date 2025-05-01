package routes

import (
	"log"
	"net/http"

	c "golang-studies/figures/controllers"
	mw "golang-studies/figures/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", c.Home)

	r.Use(mw.ContentTypeMiddleware)
	
	r.HandleFunc("/api/figures", c.GetFigures).Methods("Get")
	r.HandleFunc("/api/figures", c.CreateFigure).Methods("Post")
	
	r.HandleFunc("/api/figures/{id}", c.GetFigure).Methods("Get")
	r.HandleFunc("/api/figures/{id}", c.UpdateFigure).Methods("Put")
	r.HandleFunc("/api/figures/{id}", c.DeleteFigure).Methods("Delete")
	
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
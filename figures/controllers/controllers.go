package controllers

import (
	"encoding/json"
	"fmt"
	db "golang-studies/figures/database"
	m "golang-studies/figures/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func GetFigures(w http.ResponseWriter, r *http.Request) {
	var f []m.Figure

	db.DB.Find(&f)

	json.NewEncoder(w).Encode(f)
}

func GetFigure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for _, figure := range m.Figures {
		if strconv.Itoa(figure.Id) == id {
			json.NewEncoder(w).Encode(figure)
		}
	}
}
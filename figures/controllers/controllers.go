package controllers

import (
	"encoding/json"
	"fmt"
	db "golang-studies/figures/database"
	m "golang-studies/figures/models"
	"net/http"

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

	var f m.Figure

	db.DB.First(&f, id)

	json.NewEncoder(w).Encode(f)
}

func CreateFigure(w http.ResponseWriter, r *http.Request) {
	var f m.Figure

	json.NewDecoder(r.Body).Decode(&f)

	db.DB.Create(&f)

	json.NewEncoder(w).Encode(f)
}

func UpdateFigure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var f m.Figure

	db.DB.First(&f, id)

	json.NewDecoder(r.Body).Decode(&f)

	db.DB.Save(&f)

	json.NewEncoder(w).Encode(&f)
}

func DeleteFigure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	
	var f m.Figure

	db.DB.Delete(&f, id)

	json.NewEncoder(w).Encode(f)
}
package controllers

import (
	"encoding/json"
	"fmt"
	m "golang-studies/figures/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func GetFigures(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(m.Figures)
}
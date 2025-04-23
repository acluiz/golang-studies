package controllers

import (
	"encoding/json"
	"fmt"
	m "golang-studies/figures/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func GetFigures(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(m.Figures)
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
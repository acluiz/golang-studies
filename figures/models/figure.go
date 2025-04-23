package models

type Figure struct {
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Figures []Figure
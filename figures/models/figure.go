package models

type Figure struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Figures []Figure
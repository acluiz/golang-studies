package main

import (
	"fmt"
	db "golang-studies/figures/database"
	m "golang-studies/figures/models"
	r "golang-studies/figures/routes"
)

func main() {
	m.Figures = []m.Figure{
		{ Id: 1, Name: "lorem ipsum", Story: "lorem ipsum" },
		{ Id: 2, Name: "lorem ipsum", Story: "lorem ipsum" },
	}

	fmt.Println("runnig figures project")
	
	db.ConnectDatabase()

	r.HandleRequest()
}
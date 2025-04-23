package main

import (
	"fmt"

	m "golang-studies/figures/models"
	r "golang-studies/figures/routes"
)

func main() {
	fmt.Println("hello, world!")

	m.Figures = []m.Figure{
		{ Name: "lorem ipsum", Story: "lorem ipsum" },
		{ Name: "lorem ipsum", Story: "lorem ipsum" },
	}

	r.HandleRequest()
}
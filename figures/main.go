package main

import (
	db "golang-studies/figures/database"
	r "golang-studies/figures/routes"
)

func main() {
	db.ConnectDatabase()

	r.HandleRequest()
}
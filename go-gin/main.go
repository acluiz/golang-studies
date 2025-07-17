package main

import (
	db "golang-studies/go-gin/database"
	r "golang-studies/go-gin/routes"
)

func main() {
  db.ConnectDatabase()

  r.HandleRequests()
}
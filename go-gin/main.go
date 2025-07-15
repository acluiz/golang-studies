package main

import (
	db "golang-studies/go-gin/database"
	m "golang-studies/go-gin/models"
	r "golang-studies/go-gin/routes"
)

func main() {
  db.ConnectDatabase()

  m.Students = []m.Student{
    {Name: "Lorem", CPF: "000000", RG: "000"},
    {Name: "Ipsum", CPF: "111111", RG: "111"},
  }

  r.HandleRequests()
}
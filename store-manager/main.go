package main

import (
	"net/http"

	r "golang-studies/store-manager/routes"

	_ "github.com/lib/pq"
)

func main() {
	r.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}


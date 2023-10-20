package main

import (
	"WebApplication/routes"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", nil)
	routes.LoadRoutes()
}

package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Sneakers", "Green", 59.99, 293},
		{"Sports Cap", "Black", 14.99, 150},
		{"Denim Jeans", "Blue", 49.99, 250},
		{"Neon Green Sneakers", "Green", 59.99, 100},
		{"Summer Yellow Dress", "Yellow", 34.99, 180},
	}
	temp.ExecuteTemplate(w, "Index", products)
}

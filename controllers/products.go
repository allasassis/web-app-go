package controllers

import (
	"WebApplication/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.FindAllProducts())
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		quantityInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Error")
		}
		models.CreateNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.Delete(productId)
	http.Redirect(w, r, "/", 301)
}

func FindProductById(w http.ResponseWriter, r *http.Request) {
	product := models.FindProductById(r.URL.Query().Get("id"))
	temp.ExecuteTemplate(w, "Edit", product)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		quantityInt, err := strconv.Atoi(quantity)
		idInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error")
		}
		models.Edit(idInt, quantityInt, name, description, priceFloat)
	}

	http.Redirect(w, r, "/", 301)
}

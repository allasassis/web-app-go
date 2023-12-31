package routes

import (
	"WebApplication/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/newproduct", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/findproduct", controllers.FindProductById)
	http.HandleFunc("/update", controllers.Edit)
}

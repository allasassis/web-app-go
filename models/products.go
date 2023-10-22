package models

import (
	"WebApplication/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConnectWithDatabase()
	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
		defer db.Close()
	}
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectWithDatabase()
	insertData, err := db.Prepare("INSERT INTO products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConnectWithDatabase()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func FindProductById(id string) Product {
	db := db.ConnectWithDatabase()
	productDb, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func Edit(id, quantity int, name, description string, price float64) {
	db := db.ConnectWithDatabase()
	insertData, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, quantity, id)
	defer db.Close()
}

package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{1, "Computadora", 499.9, "16RAM", "Laptop"},
	{2, "Mouse", 9.9, "Inalambrico", "Accesorio"},
	{3, "Teclado", 19.9, "Con luz", "Accesorio"},
}

func main() {
	var new Product
	new.GetAll()
	new = Product{4, "Cargador", 49.9, "para Mac", "Accesorio"}
	new.Save()
	new.GetAll()
	product, err := GetById(4)
	if err != nil {
		fmt.Println("error")
	} else {
		printProduct(product)
	}
}

func (p Product) Save() []Product {
	Products = append(Products, p)
	return Products
}

func (p Product) GetAll() {
	var product Product
	for _, product = range Products {
		printProduct(product)
	}
}

func GetById(id int) (Product, error) {
	var product Product
	for _, product = range Products {
		if product.Id == id {
			return product, nil
		}
	}
	return product, errors.New("No hay un producto con ese id")
}

func printProduct(product Product) {
	fmt.Println("\nId: ", product.Id,
		"\nNombre: ", product.Name,
		"\nPrecio: ", product.Price,
		"\nDescripcion ", product.Description,
		"\nCategoria: ", product.Category)
}

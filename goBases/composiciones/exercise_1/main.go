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
	var nuevo Product
	nuevo.GetAll()
	nuevo = Product{4, "Cargador", 49.9, "para Mac", "Accesorio"}
	nuevo.Save()
	nuevo.GetAll()
	product, err := GetById(4)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(product)
	}
}

func (p Product) Save() []Product {
	Products = append(Products, p)
	return Products
}

func (p Product) GetAll() {
	var pr Product
	for _, pr = range Products {
		fmt.Println("\nId: ", pr.Id, "\nName: ", pr.Name, "\nPrice: ", pr.Price, "\nDescription: ", pr.Description, "\nCategory: ", pr.Category)
	}
}

func GetById(id int) (Product, error) {
	var pr Product
	for _, pr = range Products {
		if pr.Id == id {
			return pr, nil
		}
	}
	return pr, errors.New("No hay un producto con ese id")
}

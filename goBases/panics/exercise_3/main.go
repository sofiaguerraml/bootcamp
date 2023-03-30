package main

import (
	"fmt"
)

type Costumer struct {
	file    string
	name    string
	dni     int
	phone   int
	address string
}

var Costumers = []Costumer{}

var errorGeneral error

func main() {
	defer func() {
		if errorGeneral != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("Fin de la ejecucion")
	}()
	registro("doc1", "Sofia", 12345, 99, "Montevideo")
	registro("doc2", "Sofia", 12345, 99, "San Jose")
	registro("doc3", "", 123475, 99, "Colonia")
	registro("doc4", "Sofia", 7890, 19, "Maldonado")
	printAll()
}

func registro(file string, name string, dni int, phone int, address string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			errorGeneral = fmt.Errorf("Se detectaron varios errores en tiempo de ejecución")
		}
	}()
	if existe(dni) {
		panic("Error: el cliente ya existe")
	}
	new := Costumer{
		file:    file,
		name:    name,
		dni:     dni,
		phone:   phone,
		address: address,
	}
	nulo := zero(new)
	if nulo {
		panic("Error: uno o mas valores es nulo")
	}
	Costumers = append(Costumers, new)
}

func existe(dni int) bool {
	for _, c := range Costumers {
		if c.dni == dni {
			return true
		}
	}
	return false
}

func zero(new Costumer) bool {
	if new.file != "" && new.name != "" && new.dni != 0 && new.phone != 0 && new.address != "" {
		return false
	}
	return true
}

func printAll() {
	for _, costumer := range Costumers {
		fmt.Println("\nLegajo:", costumer.file,
			"\nNombre:", costumer.name,
			"\nDNI:", costumer.dni,
			"\nTelefono:", costumer.phone,
			"\nDireccion:", costumer.address)
	}
}

package main

import (
	"fmt"
)

type Costumer struct {
	legajo    string
	nombre    string
	dni       int
	telefono  int
	domicilio string
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
	fmt.Println(Costumers)

}

func registro(legajo string, nombre string, dni int, telefono int, domicilio string) {
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
	nuevo := Costumer{
		legajo:    legajo,
		nombre:    nombre,
		dni:       dni,
		telefono:  telefono,
		domicilio: domicilio,
	}
	nulo := zero(nuevo)
	if nulo {
		panic("Error: uno o mas valores es nulo")
	}
	Costumers = append(Costumers, nuevo)
}

func existe(dni int) bool {
	for _, c := range Costumers {
		if c.dni == dni {
			return true
		}
	}
	return false
}

func zero(nuevo Costumer) bool {
	if nuevo.legajo != "" && nuevo.nombre != "" && nuevo.dni != 0 && nuevo.telefono != 0 && nuevo.domicilio != "" {
		return false
	}
	return true
}

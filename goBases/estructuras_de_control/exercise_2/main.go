package main

import "fmt"

func main() {
	mensaje := otorgoPrestamo(23, true, 2, 120000)
	fmt.Println(mensaje)

}

func otorgoPrestamo(edad int, tieneEmpleo bool, antiguedad int, sueldo int) string {
	if edad > 22 && tieneEmpleo && antiguedad > 1 {
		if sueldo > 100000 {
			return "Otrogo prestamo sin intereses"
		}
		return "Otrgo prestamo con intereses"
	}
	return "No otorgo prestamo"
}

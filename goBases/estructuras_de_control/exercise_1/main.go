package main

import "fmt"

func main() {
	cantidadLetras := letrasDeUnaPalabra("Sofia")
	fmt.Println(cantidadLetras)
	imprimirLetras("Sofia")
}

func letrasDeUnaPalabra(palabra string) int {
	return len(palabra)
}

func imprimirLetras(palabra string) {
	cantidadLetras := letrasDeUnaPalabra(palabra)
	for i := 0; i < cantidadLetras; i++ {
		letra := palabra[i]
		fmt.Println(string(letra))
	}
}

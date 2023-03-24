package main

import "fmt"

func main() {
	var (
		temperatura        float64 = 24.5
		humedad            float64 = 80.1
		presionAtmosferica float64 = 1017.5
	)

	fmt.Println("Temperatura:", temperatura)
	fmt.Println("Humedad:", humedad, "%")
	fmt.Println("Presion:", presionAtmosferica)

}

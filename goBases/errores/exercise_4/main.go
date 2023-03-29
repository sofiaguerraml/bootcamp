package main

import (
	"fmt"
)

func main() {
	salary := 500
	err := esMenor(salary)
	if err != nil {
		//if errors.Is(err, errors.New("Error: el salario es menor a 10.000")) {
		fmt.Println(err)
		//}
	} else {
		fmt.Println(salary)
	}
}

func esMenor(salary int) error {
	if salary <= 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)

	}
	return nil
}

package main

import (
	"fmt"
)

func main() {
	salary := 500
	err := isLess(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}

func isLess(salary int) error {
	if salary <= 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)

	}
	return nil
}

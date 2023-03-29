package main

import (
	"errors"
	"fmt"
)

func main() {

	hours := 70.0
	cost := 50.3
	salary, err := allSalary(hours, cost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}

}

func allSalary(hours float64, cost float64) (float64, error) {
	salary := hours * cost
	if hours < 80 {
		return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if salary >= 150000 {
		return salary * 0.9, nil
	}
	return salary, nil
}

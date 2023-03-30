package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

func (m myError) Error() string {
	return m.msg
}

func main() {
	salary := 50
	e := myError{"Error: el salario es menor a 10.000"}
	err := isLess(salary)
	if err != nil {
		if errors.Is(err, e) {
			fmt.Println(e.msg)
			return
		}
	}
	fmt.Println(salary)

}

func isLess(salary int) error {
	e := myError{}
	if salary <= 10000 {
		e.msg = "Error: el salario es menor a 10.000"
		return e
	}
	return nil
}

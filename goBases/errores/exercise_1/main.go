package main

import "fmt"

type myError struct {
	msg string
}

func (m myError) Error() string {
	return m.msg
}

func main() {
	salary := 20000
	e := myError{"Error: el salario ingresado no alcanza el mínimo imponible"}
	if salary < 150000 {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

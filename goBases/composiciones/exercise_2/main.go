package main

import (
	"fmt"
	"time"
)

type Person struct {
	IdP         int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	IdE      int
	Position string
	Person
}

func main() {
	person := Person{1, "Sofia", time.Date(1997, time.November, 27, 0, 0, 0, 0, time.UTC)}
	employee := Employee{11, "Software Developer", person}
	employee.PrintEmployee()

}
func (e Employee) PrintEmployee() {
	fmt.Println(e.IdP, e.Name, e.DateOfBirth.Day(), e.DateOfBirth.Month(), e.DateOfBirth.Year(), e.IdE, e.Position)
}

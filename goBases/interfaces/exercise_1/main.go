package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name            string
	Surname         string
	DNI             int
	DateOfAdmission time.Time
}

func main() {

	var Students = []Student{
		{"Sofia", "Guerra", 123, time.Date(2022, time.March, 6, 0, 0, 0, 0, time.UTC)},
		{"Juan", "Perez", 5667, time.Date(2022, time.March, 8, 0, 0, 0, 0, time.UTC)},
		{"Luis", "Garcia", 123, time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC)},
	}

	for _, s := range Students {
		s.detail()
	}

}

func (s Student) detail() {
	fmt.Println("\nNombre: ", s.Name, "\nApellido: ", s.Surname, "\nDNI: ", s.DNI, "\nFecha: ", s.DateOfAdmission.Day(), s.DateOfAdmission.Month(), s.DateOfAdmission.Year())
}

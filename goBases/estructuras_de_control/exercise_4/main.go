package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Bejamin"])
	count := older21(employees)
	fmt.Println("Empleados mayores que 21:", count)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}

func older21(employees map[string]int) int {
	var count int = 0
	for _, element := range employees {
		if element > 21 {
			count++
		}
	}
	return count
}

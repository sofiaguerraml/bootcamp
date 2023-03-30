package main

import "fmt"

func main() {
	message := getLoan(23, true, 2, 120000)
	fmt.Println(message)

}

func getLoan(age int, haveJob bool, antique int, salary int) string {
	if age > 22 && haveJob && antique > 1 {
		if salary > 100000 {
			return "Otrogo prestamo sin intereses"
		}
		return "Otrgo prestamo con intereses"
	}
	return "No otorgo prestamo"
}

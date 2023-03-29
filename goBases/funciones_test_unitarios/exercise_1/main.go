package ej1

import "fmt"

func main() {
	var salary float64 = 150000
	tax := calculateTax(salary)
	fmt.Println(tax)

}

func calculateTax(salary float64) float64 {
	const (
		c = 50000
		m = 150000
	)
	if salary > m {
		return 0.27 * salary
	}
	if salary > c {
		return 0.17 * salary
	}
	return 0
}

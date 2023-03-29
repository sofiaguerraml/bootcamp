package ex3

import "fmt"

func main() {
	fmt.Println(calculateSalary(120, "Categoria A"))
}

func calculateSalary(minutes float64, category string) float64 {
	switch category {
	case "Categoria C":
		return 1000 * (minutes / 60)
	case "Categoria B":
		return 1500 * (minutes / 60) * 1.2
	case "Categoria A":
		return 3000 * (minutes / 60) * 1.5
	}
	return 0
}

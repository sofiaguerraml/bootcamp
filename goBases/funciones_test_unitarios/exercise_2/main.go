package ex2

import "fmt"

func main() {
	fmt.Println(average(1.5, 3, 6, 3, 7, 8, 9, 9, 10))

}

func average(grades ...float64) float64 {
	var avg float64
	for _, grade := range grades {
		avg += grade
	}
	return avg / float64(len(grades))
}

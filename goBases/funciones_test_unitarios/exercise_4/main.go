package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println(minValue)
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(averageValue)
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(maxValue)
	}

}

func operation(oper string) (func(values ...float64) float64, error) {
	switch oper {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	}
	return nil, errors.New("Error operacion")
}

func minFunc(values ...float64) float64 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func averageFunc(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func maxFunc(values ...float64) float64 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

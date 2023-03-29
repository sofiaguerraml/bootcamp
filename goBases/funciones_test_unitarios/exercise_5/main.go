package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func main() {

	animalDog, err := animal(dog)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		amount := animalDog(10)
		fmt.Println(amount)

	}

	animalCat, err := animal(cat)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		amount := animalCat(10)
		fmt.Println(amount)

	}

	animalHamster, err := animal(hamster)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		amount := animalHamster(2)
		fmt.Println(amount)

	}

	animalSpider, err := animal(spider)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		amount := animalSpider(2)
		fmt.Println(amount)

	}

}

func animal(animalType string) (func(int) float64, error) {
	switch animalType {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case spider:
		return animalSpider, nil
	}
	return nil, errors.New("No se encontro ese tipo de animal")

}

func animalDog(amount int) float64 {
	return float64(amount) * 10
}

func animalCat(amount int) float64 {
	return float64(amount) * 5
}

func animalHamster(amount int) float64 {
	return float64(amount) * 0.25
}

func animalSpider(amount int) float64 {
	return float64(amount) * 0.150
}

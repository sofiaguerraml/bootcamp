package main

import (
	"fmt"
	"os"
)

func main() {
	read()
	fmt.Println("Ejecucion finalizada")
}

func read() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	read, err := os.ReadFile("/Users/sguerra/Documents/Bootcamp/goBases/panics/exercise_2/customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Printf("%s \n", read)
}

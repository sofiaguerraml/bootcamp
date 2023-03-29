package main

import (
	"fmt"
	"os"
)

func main() {
	read()
	fmt.Println("ejecucion finalizada")
}

func read() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := os.ReadFile("costumers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

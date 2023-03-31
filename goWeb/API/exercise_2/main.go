package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string
	Surname string
}

var person Person

func main() {
	router := gin.Default()

	jsonData := `{"nombre": "Andrea","apellido": "Rivas"}`

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)

	router.POST("/saludo", func(c *gin.Context) {
		message := "Hola " + person.Name + " " + person.Surname
		c.String(http.StatusOK, message)
	})

	router.Run()
}

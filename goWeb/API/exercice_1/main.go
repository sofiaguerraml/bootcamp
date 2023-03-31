package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//---------ejercicio 1----------------
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		message := "pong"
		fmt.Fprintf(c.Writer, message)
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

	//------------Otra forma(formato json)-------
	// router := gin.Default()
	//
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{ //gin.H = map[string]any
	// 		"message": "Pong",
	// 	})
	// })
	// router.Run()

	//------------------Otra forma sin GIN(visto en vivo)--------------------------
	// router := http.NewServeMux() //Otra forma
	// router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Pong"))
	// })

	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	panic(err)
	// }

}

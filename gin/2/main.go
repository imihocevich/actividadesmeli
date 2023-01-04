package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Name     string `json:"name"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context) {
		var persona Persona

		c.BindJSON(&persona)

		saludo := persona.Name + " " + persona.Apellido
		c.String(http.StatusOK, "Hola "+saludo)
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

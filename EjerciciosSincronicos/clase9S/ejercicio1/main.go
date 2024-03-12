package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {

	router := gin.Default()

	jsonPer := `{"Nombre":"Juan","Apellido":"Perez","Edad":21,"Direccion":"Calle falsa123","Telefono":"1234456","Activo":true}`

	var p Persona
	if err := json.Unmarshal([]byte(jsonPer), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(200, p)
	})

	router.Run()

}

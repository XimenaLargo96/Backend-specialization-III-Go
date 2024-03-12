package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int `json:"id"`
	Nombre    string `json:"nombre"`
	Precio    float64 `json:"precio"`
	Stock     int `json:"stock"`
	Codigo    int `json:"codigo"`
	Publicado bool `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func main() {

	router := gin.Default()


	products,err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en la lectura del archivo")
	}



}

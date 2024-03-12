package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	
)

func main() {

	res, err := os.ReadFile("./data.csv")
	if err != nil {
		fmt.Println("Error en la lectura de nuestro archivo")
	}

	data := strings.Split(string(res), ";")

	fmt.Println(data)

	var total float64

	for i := 0; i < len(data)-1; i++ {
		var line = strings.Split(string(data[i]), ",")
		// fmt.Println(line[1], line[2])
		//fmt.Println(line)

		if i != 0 {
			precio, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				fmt.Println("El precio no se pudo parsear")
			}

			cantidad, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				fmt.Println("La cantidad no se pudo parsear")
			}
			
			totalProducto := precio * cantidad

			total += totalProducto

		}
		for i := 0; i < len(line); i++ {
			fmt.Printf("%s\t\t" , line[i])
			if i == len(line)-1 {
				fmt.Print("\n")
			}
		}
	}

fmt.Printf("\t\t %.2f", total)
}

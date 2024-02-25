package main

import "fmt"

// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
// 1.  Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// 2.  Luego, imprimir cada una de las letras.
func main() {
	palabra := "AlumnosCTD"

	var cant int = len(palabra)

	fmt.Printf("La palabra %s tiene %d letras.", palabra , cant)
	fmt.Println("\n La palabra " , palabra, " tiene " , cant , " letras.")

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
	}
}

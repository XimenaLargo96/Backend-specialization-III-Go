package main

import "fmt"

// Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
// Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente. El nombre de la nueva estudiante es Gabriela.

func main(){
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka" }

	students = append(students, "Gabriela" )

	fmt.Println(students)

	


}

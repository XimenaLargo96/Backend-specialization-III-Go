package main

import "fmt"

// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.
// Por otro lado, también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del map.


func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	mayores := 0

	for _, value := range employees {
		if value > 21 {
			mayores++
		}
	}

	fmt.Println(mayores)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
	

}


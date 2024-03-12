package main

import "fmt"


func main() {
	
	precio := calcularPrecio(5000.0, 10)
	
	fmt.Println("Ejercicio 1:",precio)

	pres := prestamo(23,true,2,100000)

	fmt.Println("Ejercicio 2:",pres)
}

/*Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/
func calcularPrecio(precio float64, descuento int) float64 {
	return precio - (precio * (float64(descuento) / 100))
}

/*Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.*/

func prestamo(age int, isEmployeed bool, yearsWorked int , salary float64) string{
	var response string
	if age > 22 && isEmployeed && yearsWorked > 1 && salary > 100000 {
		response= "Se le otorga el préstamo sin interés."
	}else if age > 22 && isEmployeed && yearsWorked > 1 {
		response= "Se le otorga el préstamo con interés."
	}else{
		response= "No se le puede otorgar un préstamo."
	}
	return response
}
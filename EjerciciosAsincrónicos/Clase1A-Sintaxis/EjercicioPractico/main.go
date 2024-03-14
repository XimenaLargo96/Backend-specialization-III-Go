package main

import "fmt"

func main()  {

	result := descuento(15500, 20)

	fmt.Println(result)

	prestamo(25,true,2,150000)

}


//Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.

func descuento(precio float32, descuento int) float32 {
	return precio - precio*(float32(descuento)/100)
}
// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.

func prestamo (age int , isEmployed bool , yearsOfService int , salary float32){
	if age > 22 && isEmployed && yearsOfService > 1{
		if salary > 100000 {
			fmt.Println("Tienes un prestamo aprobado y no se te cobrará interes")
		}else{
			fmt.Println("Tienes un prestamo aprobado y se te cobrará interes mensual")
		} 
	}else{
		fmt.Println("No podemos darte un prestamo")
	}
}
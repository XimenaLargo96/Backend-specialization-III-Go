package main

import (
	"fmt"
)

func main() {

	tax := taxOfSalary(60000)

	fmt.Println("Ejercicio 1:",tax)


	qualifications:= []int{5,3,5,-2,5,3,1,3}
	

	qualificationsAverage := average(qualifications...)

	fmt.Print("Ejercicio 2: ", qualificationsAverage)


}

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
*/

func taxOfSalary(salary float64) float64 {
	var tax float64
	if salary > 50000 && salary <= 150000 {
		tax = salary - (salary * 0.17)
	} else if salary > 150000 {
		tax = salary - (salary * 0.27)
	}
	return tax
}

/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

func average(qualifications ...int) int {
	var result int
	for _, v := range qualifications {
		if v > 0 {
			result += v
		}else{
			result = 0
			fmt.Println("No se pueden introducir calificaciones negativas")
			break
		}
	}
	return result/len(qualifications)
}
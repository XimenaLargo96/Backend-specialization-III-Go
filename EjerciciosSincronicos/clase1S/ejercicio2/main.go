package main 

import "fmt"

// A qué mes corresponde
// Realizar una aplicación que contenga una variable con el número del mes.
// 1. Según el número, imprimir el mes que corresponda en texto.
// 2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

func main(){
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Novimebre", "Diciembre"}

	mes := 3

	for i:=0 ; i <= len(meses) ; i++{
		if i == mes{
			fmt.Println(meses[i-1])
		}
	}
}
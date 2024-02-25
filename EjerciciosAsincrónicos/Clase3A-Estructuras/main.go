package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre       string
	Apellido     string
	Dni          int
	FechaIngreso time.Time
}

func main() {

	alumno := Alumno{
		Nombre: "Ximena",
		Apellido: "Largo",
		Dni: 123456,
		FechaIngreso: time.Now(),
	}

	alumno.mostrarAlumno()
}

func (a *Alumno) mostrarAlumno(){
	fmt.Print("Alumno: ", a.Nombre, a.Apellido , "\n identificacion: ", a.Dni, "\nFecha de ingreso: ", a.FechaIngreso) 
}
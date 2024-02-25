package main

import "fmt"

/*
Consigna
Una empresa de redes sociales requiere implementar una estructura usuarios con funciones
que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo
y contraseña. Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.
*/

type User struct{
	Name string
	LastName string
	Age int
	Email string
	Password string
}

func (u *User) cambiarNombre(nombre, apellido string) string{
	u.Name = nombre
	u.LastName = apellido
	return fmt.Sprintln(u)
}

func (u *User) cambiarEdad(edad int) string{
	u.Age = edad
	return fmt.Sprintln(u)
}

func (u *User) cambiarCorreo(correo string) string{
	u.Email = correo
	return fmt.Sprintln(u)
}

func (u *User) cambiarContrasenia(contrasenia string) string {
	u.Password = contrasenia
	return fmt.Sprintln(u)
}

func main()  {
	usuario := User{
		Name: "Ximena",
		LastName: "Largo",
		Age: 27,
		Email: "xime@mail.com",
		Password: "contraseña123!",
	}

	fmt.Println(usuario)
	fmt.Println("--------------------------------------------")

	var userPointer *User
	userPointer = &usuario

	fmt.Println("\nCambiar nombre:")
	fmt.Println(userPointer.cambiarNombre("Maria","Perez"))
	fmt.Println("--------------------------------------------")

	fmt.Println("\nCambiar edad:")
	fmt.Println(userPointer.cambiarEdad(25))
	fmt.Println("--------------------------------------------")

	fmt.Println("\nCambiar correo:")
	fmt.Println(userPointer.cambiarCorreo("flofy@gmail.com"))
	fmt.Println("--------------------------------------------")

	fmt.Println("\nCambiar contraseña:")
	fmt.Println(userPointer.cambiarContrasenia("contra123*!"))
	fmt.Println("--------------------------------------------")
}
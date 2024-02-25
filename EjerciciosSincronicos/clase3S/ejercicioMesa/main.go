package main

import "fmt"

/*
Ejercicio
Crear un programa que cumpla los siguientes puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el
slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá
imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto
correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().
*/
var products = []Product{
	{ID: 1, Name: "Producto1", Price: 5000.0, Description: "Este es el producto1", Category: "Categoria1"},
	{ID: 2, Name: "Producto2", Price: 1000.0, Description: "Este es el producto2", Category: "Categoria2"},
	{ID: 3, Name: "Producto3", Price: 15000.0, Description: "Este es el producto3", Category: "Categoria3"},
}

type Product struct {
	ID                          int
	Name, Description, Category string
	Price                       float32
}

func (p Product) Save(product Product) {
	products = append(products, product)
}

func (p Product) GetAll() {
	for _, product := range products {
		fmt.Printf("ID: %d, Nombre: %s, Price: %f, Description: %s, Category: %s\n", product.ID, product.Name, p.Price, product.Description, product.Category)
	}
}

func (p Product) getById(id int) {
	for _, product := range products {
		if product.ID == id {
			fmt.Printf("ID: %d, Nombre: %s, Price: %f, Description: %s, Category: %s", product.ID, product.Name, p.Price, product.Description, product.Category)
		}
	}
}

func main() {
	product1 := Product{
		ID:   4,
		Name: "Producto4",
		Price: 20000.0,
		Description: "Este es el producto4",
		Category: "Categoria4",
	}

	product1.GetAll()
	fmt.Println("----------------------")
	product1.Save(product1)
	product1.GetAll()
	fmt.Println("----------------------")
	product1.getById(4)
}

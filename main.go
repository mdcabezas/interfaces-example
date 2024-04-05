package main

import (
	"fmt"
)

// Definición de un struct llamado 'Item' para representar un artículo en la nota de venta
type Item struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

// Definición de un struct llamado 'Venta' para representar la nota de venta
type Venta struct {
	Items []Item
}

// Método para calcular el total de la nota de venta
func (v Venta) CalcularTotal() float64 {
	total := 0.0
	for _, item := range v.Items {
		total += item.Precio * float64(item.Cantidad)
	}
	return total
}

// Constructor para crear una nueva instancia de 'Venta' con algunos items predefinidos
func NuevaVenta() Venta {
	items := []Item{
		{"Camisa", 25.99, 2},
		{"Pantalón", 35.50, 1},
		{"Zapatos", 49.99, 1},
	}
	return Venta{Items: items}
}

func main() {
	// Crear una nueva venta usando el constructor
	venta := NuevaVenta()

	// Calcular el total de la venta
	total := venta.CalcularTotal()

	// Imprimir la nota de venta y el total
	fmt.Println("Nota de Venta:")
	for _, item := range venta.Items {
		fmt.Printf("- %s: $%.2f x %d\n", item.Nombre, item.Precio, item.Cantidad)
	}
	fmt.Printf("Total: $%.2f\n", total)
}

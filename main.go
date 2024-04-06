package main

import (
	"fmt"
	"math/rand"

	"github.com/mdcabezas/interfaces-example/pedidos"
)

func main() {

	items := []pedidos.Item{
		{Id: rand.Intn(10000), Nombre: "JAMÓN DE PAVO ACARAMELADO", Codigo: "105012", Precio: 34980, Cantidad: 1},
		{Id: rand.Intn(10000), Nombre: "MORTADELA LISA", Codigo: "103004", Precio: 12500, Cantidad: 3},
		{Id: rand.Intn(10000), Nombre: "COPPA ITALIANA", Codigo: "106001", Precio: 21370, Cantidad: 2},
	}

	// Nueva venta usando el constructor
	var venta pedidos.Vendedor = pedidos.NuevaVenta(items)

	// Calcula total venta
	lista := venta.ObtenerItems()
	total := venta.CalcularTotal()

	// Imprimir la nota de venta y el total
	fmt.Println("Nota de Venta:")

	for _, item := range lista {
		fmt.Printf("ID(%d) - %s %s: $%.2f x %d\n", item.Id, item.Codigo, item.Nombre, item.Precio, item.Cantidad)
	}

	fmt.Printf("Total: $%.2f\n", total)
}

package pedidos

func (v *Venta) CalcularTotal() float64 {
	total := 0.0
	for _, item := range v.Items {
		total += item.Precio * float64(item.Cantidad)
	}
	return total
}

func (v *Venta) ObtenerItems() []Item {
	return v.Items
}

func NuevaVenta(items []Item) *Venta {
	return &Venta{
		Items: items,
	}
}

package pedidos

type Item struct {
	Nombre   string
	Codigo   string
	Precio   float64
	Cantidad int
}

type Venta struct {
	Items []Item
}

type Vendedor interface {
	CalcularTotal() float64
	ObtenerItems() []Item
}

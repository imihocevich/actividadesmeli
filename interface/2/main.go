package main

import "fmt"

type Item struct {
	Nombre      string
	Description string
	Costo       float64
}

type Pequeño struct {
	Item
}

type Mediano struct {
	Item
}

type Grande struct {
	Item
}

func (p *Pequeño) Precio() float64 {
	return p.Costo
}

func (p *Mediano) Precio() float64 {
	return p.Costo + (p.Costo * 0.03)
}
func (p *Grande) Precio() float64 {
	return p.Costo + 2500 + (p.Costo * 0.06)
}

type Producto interface {
	Precio() float64
}

func Factory(nombre string, tipo string, precio float64) (producto Producto) {
	switch tipo {
	case "Grande":
		producto = &Grande{Item{nombre, tipo, precio}}
	case "Mediano":
		producto = &Mediano{Item{nombre, tipo, precio}}
	case "Pequeño":
		producto = &Pequeño{Item{nombre, tipo, precio}}
	default:
		panic("invalid tipo")
	}
	return
}

func main() {
	g := Factory("Item grande", "Grande", 1000)
	m := Factory("Item mediano", "Mediano", 900)
	p := Factory("Item pequeño", "Pequeño", 1500)

	fmt.Println(g.Precio())
	fmt.Println(m.Precio())
	fmt.Println(p.Precio())
}

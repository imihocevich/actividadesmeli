package main

const (
	grande  = "Grande"
	mediano = "Mediano"
	pequeño = "Pequeño"
)

type Item struct {
	Nombre      string
	Description string
	Costo       float64
	Tipo        string
}

type Producto interface {
	Precio() (float64, error)
}

func (i Item) Precio() (precio float64, err error) {
	switch i.Tipo {
	case "Grande":
		precio = i.Costo + 2500 + (i.Costo * 0.06)
	case "Mediano":
		precio = i.Costo + (i.Costo * 0.03)
	case "Pequeño":
		precio = i.Costo
	default:
		err = erros.New("tipo inválido")
	}
	return precio, err

}

func Factory(tipo string, precio float64) Producto {
	return
}

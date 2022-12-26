package main

import "fmt"

func main() {
	fmt.Println(Operacion(3, 3, Suma))
}

const (
	Suma           = "+"
	Resta          = "-"
	Multiplicacion = "*"
	Division       = "/"
)

func Operacion(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Suma:
		return valor1 + valor2
	case Resta:
		return valor1 - valor2
	case Multiplicacion:
		return valor1 * valor2
	case Division:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

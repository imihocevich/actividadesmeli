package main

import "fmt"

func main() {
	fmt.Println(OperacionAritmetica(Suma, 2, 3, 4, 5, 6, 7))
	fmt.Println(OperacionAritmetica(Resta, 2, 3, 4, 5, 6, 7))

}

const (
	Suma           = "+"
	Resta          = "-"
	Multiplicacion = "*"
	Division       = "/"
)

func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}
func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2

}
func opMultiplicacion(valor1, valor2 float64) float64 {
	return valor1 * valor2
}
func opDivision(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func Generador(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var resultado float64
	for _, v := range valores {
		resultado = operacion(resultado, v)
	}
	return resultado
}

func OperacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return Generador(valores, opSuma)
	case Resta:
		return Generador(valores, opResta)
	case Division:
		return Generador(valores, opDivision)
	case Multiplicacion:
		return Generador(valores, opMultiplicacion)
	}
	return 0
}

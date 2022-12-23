package main

// import "fmt"

// func main() {
// 	// c := suma(15, 15, 15)
// 	// fmt.Println(c)
// 	fmt.Println(suma(1, 1, 1))
// }

// func suma(a int, b int, c int) int {
// 	sum := a + b + c
// 	return sum
// }

// func main() {
// 	fmt.Println(Operacion(3, 3, Suma))
// }

// const (
// 	Suma           = "+"
// 	Resta          = "-"
// 	Multiplicacion = "*"
// 	Division       = "/"
// )

// func Operacion(valor1, valor2 float64, operador string) float64 {
// 	switch operador {
// 	case Suma:
// 		return valor1 + valor2
// 	case Resta:
// 		return valor1 - valor2
// 	case Multiplicacion:
// 		return valor1 * valor2
// 	case Division:
// 		if valor2 != 0 {
// 			return valor1 / valor2
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	fmt.Println("la suma de los valores es:", suma(10, 30, 40))
// }

// func suma(valores ...float64) float64 {
// 	var resultado float64
// 	for _, v := range valores {
// 		resultado += v
// 	}
// 	return resultado
// }

// func main() {
// 	fmt.Println(OperacionAritmetica(Suma, 2, 3, 4, 5, 6, 7))
// 	fmt.Println(OperacionAritmetica(Resta, 2, 3, 4, 5, 6, 7))

// }

// const (
// 	Suma           = "+"
// 	Resta          = "-"
// 	Multiplicacion = "*"
// 	Division       = "/"
// )

// func opSuma(valor1, valor2 float64) float64 {
// 	return valor1 + valor2
// }
// func opResta(valor1, valor2 float64) float64 {
// 	return valor1 - valor2

// }
// func opMultiplicacion(valor1, valor2 float64) float64 {
// 	return valor1 * valor2
// }
// func opDivision(valor1, valor2 float64) float64 {
// 	if valor2 == 0 {
// 		return 0
// 	}
// 	return valor1 / valor2
// }

// func Generador(valores []float64, operacion func(value1, value2 float64) float64) float64 {
// 	var resultado float64
// 	for _, v := range valores {
// 		resultado = operacion(resultado, v)
// 	}
// 	return resultado
// }

// func OperacionAritmetica(operador string, valores ...float64) float64 {
// 	switch operador {
// 	case Suma:
// 		return Generador(valores, opSuma)
// 	case Resta:
// 		return Generador(valores, opResta)
// 	case Division:
// 		return Generador(valores, opDivision)
// 	case Multiplicacion:
// 		return Generador(valores, opMultiplicacion)
// 	}
// 	return 0
// }

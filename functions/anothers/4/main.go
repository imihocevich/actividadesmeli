package main

import "fmt"

func main() {
	oper := OperacionAritmetica("Multiplicacion")
	r := oper(3, 6)
	fmt.Println(r)
}

func OperacionAritmetica(operador string) func(v1, v2 float64) float64 {
	switch operador {
	case "Suma":
		return opSuma
	case "Resta":
		return opResta
	case "Division":
		return opDiv
	case "Multiplicacion":
		return opMulti
	}
	return nil
}

func opSuma(v1, v2 float64) float64 {
	return v1 + v2
}
func opResta(v1, v2 float64) float64 {
	return v1 - v2
}
func opMulti(v1, v2 float64) float64 {
	return v1 * v2
}
func opDiv(v1, v2 float64) float64 {
	if v2 == 0 {
		return 0
	}

	return v1 / v2
}

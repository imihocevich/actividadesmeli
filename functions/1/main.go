package main

import "fmt"

func main() {

	fmt.Println(calculadoraSalario(30000))
}

func calculadoraSalario(salario float64) (resultado float64) {
	if salario > 150000 {
		resultado = salario * 0.27
		return
	} else if salario > 50000 {
		resultado = salario * 0.17
		return
	} else {
		return resultado
	}
}

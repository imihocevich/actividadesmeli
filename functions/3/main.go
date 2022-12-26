package main

import "fmt"

func main() {
	fmt.Printf("%.2f \n", calculadora(300000, "A"))
}

func calculadora(mins int, categoria string) (salario float64) {
	switch {
	case categoria == "C":
		salario = float64(mins/60) * 1000
	case categoria == "B":
		salario = (float64(mins/60) * 1500) * 1.2
	case categoria == "A":
		salario = (float64(mins/60) * 3000) * 1.5
	}

	if mins == 0 || mins < 0 {
		return 0
	}

	return
}

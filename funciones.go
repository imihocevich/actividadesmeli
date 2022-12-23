package main

import (
	"errors"
	"fmt"
)

// func main() {

// 	fmt.Println(calculadoraSalario(30000))
// }

// func calculadoraSalario(salario float64) (resultado float64) {
// 	if salario > 150000 {
// 		resultado = salario * 0.27
// 		return
// 	} else if salario > 50000 {
// 		resultado = salario * 0.17
// 		return
// 	} else {
// 		return resultado
// 	}
// }

// func main() {
// 	promedio, err := promedio()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(promedio)
// 	}
// }

// func promedio(notas ...float64) (resultado float64, err error) {

// 	if len(notas) == 0 {
// 		err = errors.New("No se ingresaron notas")
// 	}
// 	for _, nota := range notas {
// 		if nota < 0 {
// 			err = errors.New("No podes ingresar una nota negativa")
// 			return
// 		}
// 		resultado += nota

// 	}
// 	return resultado / float64(len(notas)), err

// }

// func main() {
// 	fmt.Printf("%.2f \n", calculadora(30000, "A"))
// }

// func calculadora(mins int, categoria string) (salario float64) {
// 	switch {
// 	case categoria == "C":
// 		salario = float64(mins/60) * 1000
// 	case categoria == "B":
// 		salario = (float64(mins/60) * 1500) * 1.2
// 	case categoria == "A":
// 		salario = (float64(mins/60) * 3000) * 1.5
// 	}
// 	return

// }

func main() {
	op, err := operacionSeleccionada(maximum)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var result = op(4, 7, 9)
	fmt.Println(result)
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func maximo(nums ...int) (resultado float64) {
	for _, numero := range nums {
		if numero > int(resultado) {
			resultado = float64(numero)
		}
	}
	return
}

func minimo(nums ...int) (resultado float64) {
	for _, numero := range nums {
		if numero < int(resultado) {
			resultado = float64(numero)
		}
	}
	return
}

func promedio(nums ...int) (resultado float64) {
	for _, numero := range nums {
		resultado += float64(numero)

	}
	return resultado / float64(len(nums))
}

func operacionSeleccionada(operador string) (operaciones func(...int) float64, err error) {
	switch operador {
	case "minimum":
		operaciones = minimo
	case "avergae":
		operaciones = promedio
	case "maximum":
		operaciones = maximo
	default:
		err = errors.New("Operacion no soportada")
	}
	return
}

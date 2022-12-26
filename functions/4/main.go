package main

import (
	"errors"
	"fmt"
)

func main() {
	op, err := operacionSeleccionada(average)

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
	for i, numero := range nums {
		if i == 0 || numero > int(resultado) {
			resultado = float64(numero)
		}
	}
	return
}

func minimo(nums ...int) (resultado float64) {
	for i, numero := range nums {
		if i == 0 || numero < int(resultado) {
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
	case "average":
		operaciones = promedio
	case "maximum":
		operaciones = maximo
	default:
		err = errors.New("operacion no soportada")
	}
	return
}

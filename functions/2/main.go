package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := promedio()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(promedio)
	}
}

func promedio(notas ...float64) (resultado float64, err error) {

	if len(notas) == 0 {
		err = errors.New("no se ingresaron notas")
	}
	for _, nota := range notas {
		if nota < 0 {
			err = errors.New("no podes ingresar una nota negativa")
			return
		}
		resultado += nota

	}
	return resultado / float64(len(notas)), err

}

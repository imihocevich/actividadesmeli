package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := salario(79, 900)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)

}

func salario(horas float64, valorHora float64) (resultado float64, err error) {
	if horas < 80 || horas < 0 {
		return 0, errors.New("no se puede trabajar menos de 80 horas")
	}

	resultado = horas * valorHora

	if resultado >= 150.000 {
		return resultado * 0.9, nil
	}

	return resultado, nil
}

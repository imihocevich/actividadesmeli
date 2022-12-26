package main

import "fmt"

func main() {
	fmt.Println("la suma de los valores es:", suma(10, 30, 40))
}

func suma(valores ...float64) float64 {
	var resultado float64
	for _, v := range valores {
		resultado += v
	}
	return resultado
}

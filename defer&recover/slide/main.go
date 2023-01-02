package main

import "fmt"

func isPair(numero int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if numero%2 != 0 {
		panic("no es un numero par")
	}
	fmt.Printf("%d es un numero par\n", numero)
}

func main() {
	numero := 5
	isPair(numero)
	fmt.Println("completada")
}

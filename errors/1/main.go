package main

import "fmt"

type customError struct {
}

func (e customError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	salary := 149000
	if salary < 150000 {
		fmt.Println(customError{})
	} else {
		fmt.Println("debe pagar impuesto")
	}
}

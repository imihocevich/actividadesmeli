package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 11000
	if salary < 10000 {
		fmt.Println(errors.New("Error: el salario es menor a 10.000"))
		return
	}
	fmt.Println("el salario es acorde	")
}

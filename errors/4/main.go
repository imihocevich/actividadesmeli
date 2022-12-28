package main

import (
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 10000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de:%d", salary)

	}
	return nil
}

func main() {
	err := checkSalary(9000)
	if err != nil {
		fmt.Println(err)
	}
}

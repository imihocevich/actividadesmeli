package main

import (
	"errors"
	"fmt"
)

type customError struct {
}

func (e customError) Error() string {
	return "Error: el salario es menor a 10.000"
}

func checkSalary(numero int) error {
	if numero < 10000 {
		fmt.Println(customError{})
	}
	return nil
}
func main() {
	salary := 8000
	err := checkSalary(salary)
	coincidence := errors.Is(err, customError{})
	fmt.Println(coincidence)
}

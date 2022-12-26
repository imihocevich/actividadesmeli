package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])
	var contador = 0

	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("Cantidad mayores a 21 años:", contador)

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}

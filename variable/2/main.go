package main

import "fmt"

func main() {
	var edad int
	var empleado bool
	var antiguedad int
	var sueldo uint

	fmt.Println("Ingrese edad")
	fmt.Scan(&edad)
	fmt.Println("Es empleado?")
	fmt.Scan(&empleado)
	fmt.Println("Ingrese antiguedad")
	fmt.Scan(&antiguedad)
	fmt.Println("Ingrese sueldo")
	fmt.Scan(&sueldo)

	if edad > 22 && empleado == true && antiguedad > 1 {
		fmt.Println("Usted esta apto para un credito")
	} else {
		fmt.Println("No puede obtener un credito")
	}

	if sueldo > 100000 {
		fmt.Println("Bonificado a tasa cero")
	}

}

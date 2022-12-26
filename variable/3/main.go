package main

import "fmt"

func main() {
	var dia int
	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo"}

	fmt.Println("Ingrese numero del mes")
	fmt.Scan(&dia)

	fmt.Println(meses[dia])

}

package main

import "fmt"

func main() {
	var word string
	fmt.Println("Ingrese la palabra a deletrear")
	fmt.Scan(&word)

	fmt.Println("La cantidad de letras que tiene son:", len(word))
	for _, i := range word {
		fmt.Println(string(i))
	}

}

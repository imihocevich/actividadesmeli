package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando ejecucion...")
	_, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Finalizando ejecucion...")
}

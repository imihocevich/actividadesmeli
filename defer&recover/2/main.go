package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando ejecucion...")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println("\nFinalizando ejecucion...")
	}()

	data, err := os.ReadFile("/Users/imihocevich/Desktop/actividad1/defer&recover/2/customers.txt")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)

}

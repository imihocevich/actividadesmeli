package main

import (
	"fmt"
	"time"
)

func procesar(i int, c chan int) {
	fmt.Println(i, "Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "Finalizado")
	c <- i
}

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go procesar(i, c)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Termino el programa", <-c)

	}
	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Termino el programa")
}

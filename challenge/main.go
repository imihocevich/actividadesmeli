package main

import (
	"fmt"
	"test/challenge/tickets"
)

func main() {
	err := tickets.ReadFile()
	if err != nil {
		panic(err)
	}

	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("El total de personas que viajaron es:", total)

	c, err := tickets.GetMornings("mañana")
	if err != nil {
		panic(err)
	}
	fmt.Println("El total de viajes en esa franja horaria es:", c)

	r, err := tickets.AverageDestination("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("El porcentaje de personas por dia es:", r)

}

package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func main() {
	Products := Product{
		ID:          1,
		Name:        "bomba",
		Price:       3.3,
		Description: "para la pileta",
		Category:    "para exteriores",
	}
	fmt.Println(Products)
}

func (p Product) Save(products ...Product) {

}

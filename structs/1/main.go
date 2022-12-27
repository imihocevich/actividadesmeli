package main

import (
	"errors"
	"fmt"
)

func main() {
	products.Getall()
	producto, err := getById(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(producto)

	producto5 := Product{
		ID:          5,
		Name:        "prueba save",
		Price:       33.3,
		Description: " ileta",
		Category:    "para ",
	}
	producto5.Save()
	products.Getall()

}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}
type Products []Product

var products = Products{
	{ID: 1,
		Name:        "bomba",
		Price:       3.3,
		Description: "para la pileta",
		Category:    "para exteriores",
	},
	{ID: 2,
		Name:        "filtro",
		Price:       9.3,
		Description: "para la pileta y baño",
		Category:    "para exteriores/interior",
	},
	{ID: 3,
		Name:        "ceramico",
		Price:       10.3,
		Description: "para la pileta",
		Category:    "para exteriores/interior",
	},
	{ID: 4,
		Name:        "barra",
		Price:       93.3,
		Description: "para lavador",
		Category:    "para exteriores/interior",
	},
}

func (p *Product) Save() {
	products = append(products, *p)
}

func (p Products) Getall() {
	for _, product := range p {
		fmt.Println(product)

	}
}

func getById(id int) (producto *Product, err error) {
	if id <= 0 {
		return producto, errors.New("invalid id")
	}

	for _, p := range products {
		if p.ID == id {
			producto = &p
			break
		}
	}

	return producto, nil
}

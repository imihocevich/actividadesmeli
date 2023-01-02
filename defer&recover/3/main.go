package main

import (
	"fmt"
)

type Cliente struct {
	Legajo    uint
	Nombre    string
	DNI       uint
	Telefono  uint
	Domicilio string
}
type Clientes []Cliente

var clientes = Clientes{
	{Legajo: 1,
		Nombre:    "João",
		DNI:       13232123213,
		Telefono:  221312312,
		Domicilio: "João",
	},
	{Legajo: 2,
		Nombre:    "vini",
		DNI:       1123123,
		Telefono:  287678,
		Domicilio: "vini",
	},
	{
		Legajo:    3,
		Nombre:    "david",
		DNI:       19889798,
		Telefono:  2871266781,
		Domicilio: "david",
	},
	{
		Legajo:    0,
		Nombre:    "david",
		DNI:       19889798,
		Telefono:  2871266781,
		Domicilio: "david",
	},
}

func isActive(legajo uint) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	for _, v := range clientes {
		if v.Legajo == legajo {
			panic("Error: el cliente esta activado")
		}

	}
	fmt.Printf("%v no esta registrado", legajo)
}

//	func isNotZero() (output string, err error) {
//		for _, cliente := range clientes {
//			if cliente.Legajo == 0 || cliente.Nombre == "" || cliente.Domicilio == "" || cliente.Telefono == 0 || cliente.DNI == 0 {
//				return "", errors.New("deberia ingresar campos correctamente")
//			}
//		}
//		return "cliente con valores validos", nil
//	}
func main() {
	isActive(2)
}

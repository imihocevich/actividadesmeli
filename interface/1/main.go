package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    int
}

func (alumnos *Alumnos) Alldata() string {
	return fmt.Sprintf("%s %s %d %d,\n", alumnos.Nombre, alumnos.Apellido, alumnos.DNI, alumnos.Fecha)
}

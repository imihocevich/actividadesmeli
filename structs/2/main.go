package main

import "fmt"

func main() {
	employee1 := Employee{
		ID:       1,
		Position: "Junior",
		Person: Person{
			ID:          2,
			Name:        "Ivo",
			DateOfBirth: 14091994},
	}

	fmt.Println(employee1)
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth int
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("%d %s %d %s %d,\n", e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth)
}

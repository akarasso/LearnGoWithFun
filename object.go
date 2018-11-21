package main

import (
	"fmt"
)

type Person struct {
	Name string
}

/*
**	Definition d'un objet etenduw
*/

type Android struct {
	Person
	Model string
}

func (p *Person) Talk()  {
	fmt.Println("Heyyyy", p.Name, "!!")
}

func main() {
	alex := Person{Name: "Alex"}
	darwin := Android{Person:Person{Name: "Darwin"}}
	alex.Talk()
	darwin.Person.Talk()
}

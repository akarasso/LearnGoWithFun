package main

import "fmt"

type Talk interface {
	Talk()
}

type Person struct{
	Name string
}

type Android struct{
	Name string
	Age int
}

func (p Person) Talk() {
	fmt.Println("My name is", p.Name)
}

func (p Android) Talk() {
	fmt.Println("My name is", p.Name, "and I'm", p.Age)
}

func makeItTalk(g Talk) {
	g.Talk()
}

func main() {
	alex := Person{Name:"Alex"}
	darwin := Android{Name:"Darwin",Age:10}
	makeItTalk(alex)
	makeItTalk(darwin)
}

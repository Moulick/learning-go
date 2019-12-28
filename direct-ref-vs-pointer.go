package main

import "fmt"

type contact struct {
	phone int
}

type person struct {
	firstName  string
	secondName string
	contact
}

func main() {

	jim := person{
		firstName:  "jim",
		secondName: "choo",
		contact: contact{

			phone: 98788,
		},
	}

	jim.print()
	jim = jim.changeName("jimmy")
	jim.print()

}

func (p person) changeName(newName string) person {

	p.firstName = newName

	return p
}

func (p person) print() {

	fmt.Printlnp)
}

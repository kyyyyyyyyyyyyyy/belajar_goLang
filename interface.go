package main

import "fmt"

type hasName interface {
	GetName() string
}

func sayHello(value hasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	Person := Person{Name: "kayu"}
	sayHello(Person)

	animal := Animal{Name: "kucing"}
	sayHello(animal)
}

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello ", name, "my name is ", customer.Name)
}

func main() {
	var kayu Customer
	kayu.Name = "muhammad Rizky"
	kayu.Address = "kuningan"
	kayu.Age = 18

	fmt.Println(kayu)
	fmt.Println(kayu.Name)
	fmt.Println(kayu.Address)
	fmt.Println(kayu.Age)

	jaka := Customer{
		Name:    "jaka",
		Address: "solo",
		Age:     19,
	}

	fmt.Println(jaka)

	budi := Customer{"budi", "bekasi", 29}

	fmt.Println(budi)

	budi.sayHello("agus")
	jaka.sayHello("agus")
}

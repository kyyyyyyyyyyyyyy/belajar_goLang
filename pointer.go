package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"kuningan", "jawa barat", "indonesia"}
	// address2 := address1

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"kuningan", "jawa barat", "indonesia"}
	address2 := & address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "indonesia"
}

func main() {
	address := Address{}
	changeCountryToIndonesia(&address)

	fmt.Println(address)
}

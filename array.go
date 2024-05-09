package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "kayu"
	names[1] = "bakar"
	names[2] = "ireng"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		60,
		70,
		50,
		40,
	}

	fmt.Println(values)  

}

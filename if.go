package main

import "fmt"

func main() {
	name := "kayuuu"

	if name == "kayu" {
		fmt.Println("hello world")
	} else if name == "kay" {
		fmt.Println("hello kyu")
	} else {
		fmt.Println("siapa nama mu")
	}

	if length := len(name); length > 5 {
		fmt.Println("kepanjangan")
	} else {
		fmt.Println("right")
	}
}

package main

import "fmt"

func main() {
	name := "kyuuuuu"

	switch name {
	case "kayu":
		fmt.Println("hello kayu")
	case "kyuu":
		fmt.Println("hello kyu")
	default:
		fmt.Println("boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("kepanjangan")
	case false:
		fmt.Println("right")
	}
}

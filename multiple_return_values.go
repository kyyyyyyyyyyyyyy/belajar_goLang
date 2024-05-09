package main

import "fmt"

func getFullName() (string, string) {
	return "muhammad", "rizky"
}

func main() {
	// firstName, SecondName := getFullName()
	// fmt.Println(firstName, SecondName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}

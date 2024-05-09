package main

import "fmt"

func getCompleteName() (firstName string, SecondName string, Thredname string) {
	firstName = "Muhammad"
	SecondName = "Rizky"
	Thredname = "kyuu"

	return firstName, SecondName, Thredname
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

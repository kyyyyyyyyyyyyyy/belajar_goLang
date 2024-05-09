package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	misal := getGoodBye

	fmt.Println(misal("kayu"))
}

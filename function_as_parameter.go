package main

import "fmt"

//type decraation
type filter func(string) string

func sayHelloWithFilter(name string, filter filter) {
	filterdName := filter(name)
	fmt.Println("hello ", filterdName)
}

func spamFilter(name string) string {
	if name == "asu" {
		return "..."
	} else {
		return name
	}
}

func main() {
	print := sayHelloWithFilter
	filter := spamFilter
	print("kayu", filter)
	print("asu", filter)
}

package main

import "fmt"

type blacklist func(string) bool

func userRegister(name string, blacklist blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "asu"
	}

	userRegister("kayu", blacklist)

	userRegister("asu", func(name string) bool {
		return name == "asu"
	})
}

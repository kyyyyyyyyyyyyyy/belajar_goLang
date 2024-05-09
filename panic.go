package main

import "fmt"

func endApplicaition() {
	fmt.Println("end app")

	message := recover()
	fmt.Println("terjadi error", message)
}

func runApp(errno bool) {
	defer endApplicaition()

	if errno {
		panic("error")
	}
}

func main() {
	runApp(true)
	fmt.Println("kayu")
}

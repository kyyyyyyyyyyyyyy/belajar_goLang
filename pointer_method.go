package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	kayu := Man{"kayu"}
	kayu.Married()

	fmt.Println(kayu.Name)
}

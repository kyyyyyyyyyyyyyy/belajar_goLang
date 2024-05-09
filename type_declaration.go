package main

import "fmt"

func main() {
	type NIS string

	var rizky NIS = "11111111"

	var contoh string = "2222222"
	var contohNis NIS = NIS(contoh)

	fmt.Println(rizky)
	fmt.Println(contohNis)
}
